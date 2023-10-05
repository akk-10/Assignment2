package data

import (
	"time"
)

type Camera struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Model      string    `json:"model"`
	Resolution string    `json:"resolution"`
	Weight     float64   `json:"weight"`
	Zoom       float64   `json:"zoom"`
	Version    int32     `json:"version"`
}
