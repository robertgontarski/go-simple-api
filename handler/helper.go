package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
	Msg        any `json:"msg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewApiError(statusCode int, err error) ApiError {
	return ApiError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) ApiError {
	return ApiError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() ApiError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandleFunc(h ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var apiErr ApiError
			if errors.As(err, &apiErr) {
				if err := writeJSON(w, apiErr.StatusCode, apiErr); err != nil {
					slog.Error("Api error", "err", err, "path", r.URL.Path)
					return
				}
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal server error",
				}

				if err := writeJSON(w, http.StatusInternalServerError, errResp); err != nil {
					slog.Error("Server error", "err", err, "path", r.URL.Path)
					return
				}
			}

			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func AsJSONContent(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}
