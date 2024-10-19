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