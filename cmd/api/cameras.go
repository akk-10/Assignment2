package main

import (
	"fmt"
	"mycameraapp/internal/data"
	"mycameraapp/internal/validator"
	"net/http"
	"time"
)

func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string  `json:"name"`
		Model      string  `json:"model"`
		Resolution string  `json:"resolution"`
		Weight     float64 `json:"weight"`
		Zoom       float64 `json:"zoom"`
	}
	// BODY='{"Name":"Film Camera","Model":"Nikon FM2","Resolution":"35mm", "Weight":650.0, "Zoom":0.0]}'
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	Camera := &data.Camera{
		Name:       input.Name,
		Model:      input.Model,
		Resolution: input.Resolution,
		Weight:     input.Weight,
		Zoom:       input.Zoom,
		Version:    1, // Set the version to a default value
	}

	v := validator.New()
	if data.ValidateCamera(v, Camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(input.Model != "", "model", "must be provided")
	v.Check(input.Resolution != "", "resolution", "must be provided")
	v.Check(input.Weight > 0, "weight", "must be greater than 0")
	v.Check(input.Zoom >= 0, "zoom", "must be greater than or equal to 0")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	camera := data.Camera{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "Vintage Camera",
		Model:      "Canon EOS 5D Mark IV",
		Resolution: "4K",
		Weight:     800.0,
		Zoom:       5.0,
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"Camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
