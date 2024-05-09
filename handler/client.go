package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-api/models"
	"simple-api/types"
	"simple-api/utils"
	"strconv"
)

type ClientHandler struct {
	*DefaultHandler
}

func NewClientHandler(dh *DefaultHandler) *ClientHandler {
	return &ClientHandler{
		DefaultHandler: dh,
	}
}

func (h *ClientHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) error {
	var req types.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	if err := h.validate.Struct(req); err != nil {
		errs := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			key, val := utils.GetErrorMessage(err, req)
			errs[key] = val
		}

		return InvalidRequestData(errs)
	}

	client, err := models.NewClient(req.Email, req.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, &types.DefaultResponse{
			Status: http.StatusInternalServerError,
			Msg:    "error while encoding password",
			Data:   nil,
		})
	}

	if err := h.db.AddClient(client); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, &types.DefaultResponse{
		Status: http.StatusOK,
		Msg:    "client added",
		Data: map[string]any{
			"client_id": client.ID,
		},
	})
}

func (h *ClientHandler) GetClientByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return InvalidRequestData(map[string]string{
			"id": "must be an int",
		})
	}

	client, err := h.db.GetClientByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewApiError(http.StatusNotFound, fmt.Errorf("client with id: %d, not found", id))
		}
		return err
	}

	return writeJSON(w, http.StatusOK, &types.DefaultResponse{
		Status: http.StatusOK,
		Msg:    "return client",
		Data: map[string]any{
			"client": client,
		},
	})
}
