package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods("GET")
	router.HandleFunc("/v1/cameras", app.listCamerasHandler).Methods("GET")
	router.HandleFunc("/v1/cameras", app.createCameraHandler).Methods("POST")
	router.HandleFunc("/v1/cameras/{id:[0-9]+}", app.getCameraHandler).Methods("GET")
	router.HandleFunc("/v1/cameras/{id:[0-9]+}", app.updateCameraHandler).Methods("PUT")
	router.HandleFunc("/v1/cameras/{id:[0-9]+}", app.deleteCameraHandler).Methods("DELETE")

	return router
}
