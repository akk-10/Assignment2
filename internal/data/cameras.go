package data

import (
	"database/sql"
	"errors"
	"mycameraapp/internal/validator"
	"time"
)

type CameraModel struct {
	DB *sql.DB
}

func (c CameraModel) Insert(camera *Camera) error {
	query := `INSERT INTO camera (Name, Model, Resolution, Weight, Zoom) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, version`

	args := []interface{}{camera.Name, camera.Model, camera.Resolution, camera.Weight, camera.Zoom}

	return c.DB.QueryRow(query, args...).Scan(&camera.ID, &camera.CreatedAt, &camera.Version)
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Camera CameraModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Camera: CameraModel{DB: db},
	}
}

type Camera struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	Resolution string    `json:"resolution"`
	Weight     float64   `json:"weight"`
	Zoom       float64   `json:"zoom"`
	Version    int32     `json:"version"`
}

func ValidateCamera(v *validator.Validator, camera *Camera) {
	v.Check(camera.Name != "", "name", "must be provided")
	v.Check(len(camera.Name) <= 500, "name", "must not exceed 500 characters")
	v.Check(camera.Model != "", "model", "must be provided")
	v.Check(camera.Resolution != "", "resolution", "must be provided")
	v.Check(camera.Weight > 0, "weight", "must be greater than 0")
	v.Check(camera.Zoom >= 0, "zoom", "must be greater than or equal to 0")

}

func (c CameraModel) Get(id int64) (*Camera, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `SELECT id, created_at, name, model, resolution, weight, zoom FROM cameras WHERE Id = $1`
	var camera Camera

	err := c.DB.QueryRow(query, id).Scan(
		&camera.ID,
		&camera.CreatedAt,
		&camera.Name,
		&camera.Model,
		&camera.Resolution,
		&camera.Weight,
		&camera.Zoom,
		&camera.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &camera, nil
}
func (c CameraModel) Update(camera *Camera) error {
	query := `UPDATE cameras SET name = $1, model = $2, resolution = $3, weight = $4, weight = $5, version = version + 1 WHERE id = $6 RETURNING version`

	args := []interface{}{
		camera.Name,
		camera.Model,
		camera.Resolution,
		camera.Weight,
		camera.Zoom,
		camera.ID,
	}
	return c.DB.QueryRow(query, args...).Scan(&camera.Version)
}
func (c CameraModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM cameras WHERE id = $1`

	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
