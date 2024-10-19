package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"database/sql"
	"device-service/internal/domains/entities"

	_ "github.com/lib/pq"
)

const (
	connection = "user=dev dbname=device sslmode=disable password="
)

const (
	createTablesScript = `
		CREATE TABLE houses (
			id SERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			address VARCHAR(255) NOT NULL
		);
		
		CREATE TABLE devices (
			id SERIAL PRIMARY KEY,
			type_id BIGINT NOT NULL,
			house_id BIGINT NOT NULL,
			serial_number VARCHAR(255) NOT NULL,
			status VARCHAR(255) NOT NULL,
			FOREIGN KEY (house_id) REFERENCES houses(id)
		);
		
		CREATE TABLE device_types (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);
		
		CREATE TABLE modules (
			id SERIAL PRIMARY KEY,
			device_id BIGINT NOT NULL,
			name VARCHAR(255) NOT NULL,
			FOREIGN KEY (device_id) REFERENCES devices(id)
		);
		
		CREATE TABLE telemetry_data (
			id SERIAL PRIMARY KEY,
			device_id BIGINT NOT NULL,
			data TEXT NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (device_id) REFERENCES devices(id)
		);
	`
)

type repoDevice struct {
	db *sql.DB
}

func NewDeviceRepository() IDeviceRepository {
	db, err := sql.Open("postgres", os.Getenv("DEVICE_POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
	instance := &repoDevice{
		db: db,
	}
	if err := instance.init(); err != nil {
		log.Fatal(err)
	}
	return instance
}

func (r *repoDevice) GetDeviceInfo(id string) (*entities.Device, error) {
	var device entities.Device

	// Query the devices table
	err := r.db.QueryRow("SELECT id, type_id, house_id, serial_number, status FROM devices WHERE id=$1", id).
		Scan(&device.ID, &device.TypeID, &device.HouseID, &device.SerialNumber, &device.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewDeviceNotFoundError(id)
		}
		return nil, err
	}

	// Query the device_types table
	err = r.db.QueryRow("SELECT id, name FROM device_types WHERE id=$1", device.TypeID).
		Scan(&device.DeviceType.ID, &device.DeviceType.Name)
	if err != nil {
		return nil, err
	}

	// Query the houses table
	err = r.db.QueryRow("SELECT id, user_id, address FROM houses WHERE id=$1", device.HouseID).
		Scan(&device.HouseID, &device.House.UserID, &device.House.Address)
	if err != nil {
		return nil, err
	}

	// Query the modules table
	rows, err := r.db.Query("SELECT id, device_id, name FROM modules WHERE device_id=$1", device.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var module entities.Module
		err = rows.Scan(&module.ID, &module.DeviceID, &module.Name)
		if err != nil {
			return nil, err
		}
		device.Modules = append(device.Modules, module)
	}

	return &device, nil
}

func (r *repoDevice) UpdateDeviceStatus(id string, status string) error {
	result, err := r.db.Exec("UPDATE devices SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return NewDeviceNotFoundError(id)
	}
	return nil
}

func (r *repoDevice) init() error {
	_, err := r.db.Exec(createTablesScript)
	return err
}
