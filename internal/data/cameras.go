package data

import (
	"mycameraapp/internal/validator"
	"time"
)

type Camera struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	Resolution string    `json:"resolution"`
	Weight     float64   `json:"weight"`
	Runtime    Runtime   `json:"runtime,omitempty"`
	Zoom       float64   `json:"zoom"`
	Version    int32     `json:"version"`
}

func ValidateCamera(v *validator.Validator, camera *Camera) {
	v.Check(camera.Name != "", "name", "must be provided")
	v.Check(len(camera.Name) <= 500, "name", "must not exceed 500 characters")
	v.Check(camera.Model != "", "model", "must be provided")
	v.Check(camera.Resolution != "", "resolution", "must be provided")
	v.Check(camera.Weight > 0, "weight", "must be greater than 0")
	v.Check(camera.Runtime != 0, "runtime", "must be provided")
	v.Check(camera.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(camera.Zoom >= 0, "zoom", "must be greater than or equal to 0")
}
