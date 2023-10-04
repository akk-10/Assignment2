package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Create a new vintage camera item and return a placeholder response.
func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new vintage camera item")
}

// Show the details of a vintage camera item by ID.
func (app *application) getCameraHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Show the details of vintage camera item with ID %d\n", id)
}

// List all vintage camera items.
func (app *application) listCamerasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List all vintage camera items")
}

// Update the details of a vintage camera item by ID.
func (app *application) updateCameraHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Update vintage camera item with ID %d\n", id)
}

// Delete a vintage camera item by ID.
func (app *application) deleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Delete vintage camera item with ID %d\n", id)
}
