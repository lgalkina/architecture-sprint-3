package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"telemetry-service/internal/domains/entities"

	_ "github.com/lib/pq"
)

const (
	createTableScript = `CREATE TABLE telemetries (
		id SERIAL PRIMARY KEY,
		device_id VARCHAR(255) NOT NULL,
		temperature DOUBLE PRECISION NOT NULL,
		timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
)

// singleton instance of the repository
var (
	instance *repoTelemetry
	once     sync.Once
)

type repoTelemetry struct {
	db *sql.DB
}

func NewTelemetryRepository() ITelemetryRepository {
	once.Do(func() {
		db, err := sql.Open("postgres", os.Getenv("TELEMETRY_POSTGRES_URL"))
		if err != nil {
			log.Fatal(err)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully connected to the database")
		instance = &repoTelemetry{
			db: db,
		}
		if err := instance.init(); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}

func (t *repoTelemetry) GetDeviceTelemetry(id string) ([]entities.TelemetryData, error) {
	query := "SELECT device_id, temperature, timestamp FROM telemetries WHERE device_id=$1 ORDER BY timestamp DESC"
	rows, err := t.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var telemetries []entities.TelemetryData
	for rows.Next() {
		var td entities.TelemetryData
		err = rows.Scan(&td.DeviceID, &td.Temperature, &td.Timestamp)
		if err != nil {
			return nil, err
		}
		telemetries = append(telemetries, td)
	}

	return telemetries, nil
}

func (t *repoTelemetry) GetLatestDeviceTelemetry(id string) (*entities.TelemetryData, error) {
	query := "SELECT id, device_id, temperature, timestamp FROM telemetries WHERE device_id=$1 ORDER BY timestamp DESC LIMIT 1"
	var td entities.TelemetryData
	err := t.db.QueryRow(query, id).Scan(&td.DeviceID, &td.Temperature, &td.Timestamp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewDeviceTelemetryNotFoundError(id)
		}
		return nil, err
	}

	return &td, nil
}

func (t *repoTelemetry) SaveDeviceTelemetry(data *entities.TelemetryData) error {
	query := "INSERT INTO telemetries (device_id, temperature, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := t.db.Exec(query, data.DeviceID, data.Temperature, data.Timestamp)
	if err != nil {
		return err
	}

	return nil
}

func (t *repoTelemetry) init() error {
	_, err := t.db.Exec(createTableScript)
	return err
}
