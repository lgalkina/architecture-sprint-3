-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS houses (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    type_id BIGINT NOT NULL,
    house_id BIGINT NOT NULL,
    serial_number VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    FOREIGN KEY (house_id) REFERENCES houses(id)
);

CREATE TABLE IF NOT EXISTS device_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS modules (
    id SERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    FOREIGN KEY (device_id) REFERENCES devices(id)
);

CREATE TABLE IF NOT EXISTS telemetry_data (
    id SERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    data TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (device_id) REFERENCES devices(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd