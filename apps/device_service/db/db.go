package db

import (
	"context"
	"device_service/models"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents the database connection
type DB struct {
	Pool *pgxpool.Pool
}

// New creates a new DB instance
func New(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{Pool: pool}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// GetDevices retrieves all devices from the database for user
func (db *DB) GetDevices(ctx context.Context, user_id int, location_id int, type_code string) ([]models.Device, error) {
	query := `
		SELECT id, name, type, status, serial_number, location_id, model_id, last_updated, created_at
		FROM device
		WHERE owner_id = $1
	`
	args := []interface{}{user_id}
	paramIdx := 2
	// Добавление фильтра по location_id если задан
	if location_id > 0 {
		query += fmt.Sprintf(" AND location_id = $%d", paramIdx)
		args = append(args, location_id)
		paramIdx++
	}

	if type_code != "" {
		query += fmt.Sprintf(" AND type = $%d", paramIdx)
		args = append(args, type_code)
		paramIdx++
	}

	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying devices: %w", err)
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var d models.Device
		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Type,
			&d.Status,
			&d.Serial_Number,
			&d.LocationID,
			&d.DeviceModelID,
			&d.LastUpdated,
			&d.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning device row: %w", err)
		}
		devices = append(devices, d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating device rows: %w", err)
	}

	return devices, nil
}

func (db *DB) CreateDevice(ctx context.Context, d models.DeviceCreate) (models.Device, error) {
	query := `
		INSERT INTO device (name, location_id, type, model_id, serial_number, owner_id, status)
		VALUES ($1, $2, $3, $4, $5, $6, 'inactive')
		RETURNING id, name, location_id, type, model_id, serial_number, status, last_updated, created_at
	`
	// now = time.Now()
	var device models.Device
	err := db.Pool.QueryRow(ctx, query,
		d.Name,
		d.LocationID,
		d.Type,
		d.DeviceModelID,
		d.Serial_Number,
		d.OwnerID,
	).Scan(
		&device.ID,
		&device.Name,
		&device.LocationID,
		&device.Type,
		&device.DeviceModelID,
		&device.Serial_Number,
		&device.Status,
		&device.LastUpdated,
		&device.CreatedAt,
	)
	if err != nil {
		return models.Device{}, fmt.Errorf("error creating device: %w", err)
	}
	return device, nil

}
