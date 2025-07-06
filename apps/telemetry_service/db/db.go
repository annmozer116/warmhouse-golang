package db

import (
	"context"
	"fmt"
	"telemetry_service/models"

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

func (db *DB) CreateDevice(ctx context.Context, d models.Device) (models.Device, error) {
	query := `
		INSERT INTO device (id, type)
		VALUES ($1, $2)
		RETURNING id, type
	`
	var device models.Device
	err := db.Pool.QueryRow(ctx, query,
		d.DeviceID,
		d.DeviceType,
	).Scan(
		&device.DeviceID,
		&device.DeviceType,
	)
	if err != nil {
		return models.Device{}, fmt.Errorf("error creating device: %w", err)
	}
	return device, nil
}
