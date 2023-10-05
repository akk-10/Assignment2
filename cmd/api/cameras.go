package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mycameraapp/internal/data"
	"net/http"
	"strconv"
	"time"
)

func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new camera")
}
func (app *application) showCameraHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	camera := data.Camera{
		ID:         id,
		CreatedAt:  time.Now(),
		Model:      "Canon EOS 5D Mark IV",
		Resolution: "4K",
		Weight:     800.0,
		Zoom:       5.0,
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, camera, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
