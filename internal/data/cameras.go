package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"mycameraapp/internal/validator"
	"time"
)

type CameraModel struct {
	DB *sql.DB
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func (c CameraModel) Insert(camera *Camera) error {
	query := `INSERT INTO cameras (Name, Model, Resolution, Weight, Zoom) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, version`

	args := []interface{}{camera.Name, camera.Model, camera.Resolution, camera.Weight, camera.Zoom}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.DB.QueryRowContext(ctx, query, args...).Scan(&camera.ID, &camera.CreatedAt, &camera.Version)
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
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

	query := `SELECT  id, created_at, name, model, resolution, weight, zoom FROM cameras WHERE Id = $1`
	var camera Camera

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
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
	query := `UPDATE cameras SET name = $1, model = $2, resolution = $3, weight = $4, weight = $5, version = version + 1 WHERE id = $6 AND version = $7 RETURNING version`

	args := []interface{}{
		camera.Name,
		camera.Model,
		camera.Resolution,
		camera.Weight,
		camera.Zoom,
		camera.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&camera.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}
func (c CameraModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM cameras WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := c.DB.ExecContext(ctx, query, id)
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
func (c CameraModel) GetAll(name string, model string, filters Filters) ([]*Camera, Metadata, error) {
	query := fmt.Sprintf(` SELECT id, created_at, name, model, resolution, weight, zoom FROM cameras 
        WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
        AND (model @> $2 OR $2 = '') ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{name, model, filters.limit(), filters.offset()}

	rows, err := c.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	///	rows, err := c.DB.QueryContext(ctx, query, name, model)
	///	if err != nil {
	///		return nil, Metadata{}, err
	///	}

	defer rows.Close()

	totalRecords := 0
	cameras := []*Camera{}

	for rows.Next() {
		var camera Camera
		err := rows.Scan(
			&totalRecords,
			&camera.ID,
			&camera.CreatedAt,
			&camera.Name,
			&camera.Model,
			&camera.Resolution,
			&camera.Zoom,
			&camera.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		cameras = append(cameras, &camera)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return cameras, metadata, nil
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

// curl -w '\nTime: %{time_total}s \n' localhost:4000/v1/cameras/1
