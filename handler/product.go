package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-api/models"
	"simple-api/storage"
	"simple-api/types"
	"simple-api/utils"
	"strconv"
)

type ProductHandler struct {
	db       storage.Storage
	validate *validator.Validate
}

func NewProductHandler(db storage.Storage, validate *validator.Validate) *ProductHandler {
	return &ProductHandler{
		db:       db,
		validate: validate,
	}
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) error {
	var req types.CreateProductRequest
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

	product := models.NewProduct(req.Name, req.Price)
	if err := h.db.AddProduct(product); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return InvalidRequestData(map[string]string{
			"id": "must be an int",
		})
	}

	product, err := h.db.GetProductByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewApiError(http.StatusNotFound, fmt.Errorf("product with id: %d, not found", id))
		}
		return err
	}

	return writeJSON(w, http.StatusOK, product)
}
