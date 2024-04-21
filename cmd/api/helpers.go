package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

type envelope map[string]any

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func getIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParamFromCtx(r.Context(), "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return 0, errors.New("bad id param")
	}
	return id, nil
}
