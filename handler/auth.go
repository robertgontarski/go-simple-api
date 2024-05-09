package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simple-api/types"
	"simple-api/utils"
)

type AuthHandler struct {
	*DefaultHandler
}

func NewAuthHandler(dh *DefaultHandler) *AuthHandler {
	return &AuthHandler{
		DefaultHandler: dh,
	}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	var req types.LoginAuthRequest
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

	client, err := h.db.GetClientByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewApiError(http.StatusNotFound, errors.New("client not found"))
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(req.Password)); err != nil {
		return writeJSON(w, http.StatusUnauthorized, &types.DefaultResponse{
			Status: http.StatusUnauthorized,
			Msg:    "invalid cred",
			Data:   nil,
		})
	}

	token, err := h.jwtAuth.CreateToken(client)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, &types.DefaultResponse{
		Status: http.StatusOK,
		Msg:    "return token",
		Data: map[string]any{
			"token": token,
		},
	})
}
