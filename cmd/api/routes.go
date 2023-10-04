package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

// Creating a new camera item and return a placeholder response.
func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new vintage camera item")
}

// Showing the details of a vintage camera item by ID.
func (app *application) getCameraHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Show the details of vintage camera item with ID %d\n", id)
}
