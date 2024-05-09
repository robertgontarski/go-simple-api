package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"simple-api/auth"
	"simple-api/storage"
	"simple-api/types"
	"strconv"
	"time"
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

func AsProtected(handler http.HandlerFunc, db storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			writeJSON(w, http.StatusUnauthorized, &types.DefaultResponse{
				Status: http.StatusUnauthorized,
				Msg:    "missing token",
			})
		}

		tokenString = tokenString[len("Bearer "):]
		token, err := (auth.NewJWTAuth()).VerifyToken(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id, err := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
		if err != nil {
			permissionDenied(w)
			return
		}

		if _, err := db.GetClientByID(id); err != nil {
			permissionDenied(w)
			return
		}

		t, err := claims.GetExpirationTime()
		if err != nil {
			permissionDenied(w)
			return
		}

		if time.Since(t.Time) > 0 {
			permissionDenied(w)
			return
		}

		handler(w, r)
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func permissionDenied(w http.ResponseWriter) error {
	return writeJSON(w, http.StatusForbidden, NewApiError(http.StatusForbidden, errors.New("permission denied")))
}
