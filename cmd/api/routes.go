package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc(http.MethodPost, "/v1/cameras", app.createCameraHandler)
	router.HandleFunc(http.MethodGet, "/v1/cameras/:id", app.showCameraHandler)
	return router
}
