package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/ntnghiatn/financial-app-backend/internal/api/v1"
	"github.com/ntnghiatn/financial-app-backend/internal/database"
)

func NewRouter(db database.Database) (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
