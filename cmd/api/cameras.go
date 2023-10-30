package main

import (
	"errors"
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
	// BODY='{"Name":"Film Camera","Model":"Nikon FM2","Resolution":"35mm", "Weight":650.0"]}'
	// BODY='{"Name":"Film Camera","Model":"Nikon FM2","Resolution":"35mm","Weight":650.0}'
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
		Version:    1,
	}

	v := validator.New()
	if data.ValidateCamera(v, Camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Camera.Insert(Camera)
	if err != nil {
		app.serverErrorResponse(w, r, err)
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
func (app *application) updateCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	camera, err := app.models.Camera.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Name       *string  `json:"name"`
		Model      *string  `json:"model"`
		Resolution *string  `json:"resolution"`
		Weight     *float64 `json:"weight"`
		Zoom       *float64 `json:"zoom"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//camera.Name = input.Name
	//camera.Model = input.Model
	//camera.Resolution = input.Resolution
	//camera.Weight = input.Weight
	//camera.Zoom = input.Zoom

	if input.Name != nil {
		camera.Name = *input.Name
	}
	if input.Model != nil {
		camera.Model = *input.Model
	}
	if input.Resolution != nil {
		camera.Resolution = *input.Resolution
	}
	if input.Weight != nil {
		camera.Weight = *input.Weight
	}
	if input.Zoom != nil {
		camera.Zoom = *input.Zoom
	}

	v := validator.New()
	if data.ValidateCamera(v, camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Camera.Update(camera)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Camera.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// BODY='{"Name":"Film Camera","Model":"Nikon FM2","Resolution":"35mm", "Weight":650.0"}'
//curl -X PUT -d "$BODY" localhost:4000/v1/cameras/2
// curl -X PATCH -d '{"weight: 750.0}' localhost:4000/v1/cameras/1
