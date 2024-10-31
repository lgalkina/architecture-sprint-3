-- +goose Up
-- +goose StatementBegin
INSERT INTO houses (user_id, address) VALUES
    (1, '123 Elm Street'),
    (2, '456 Oak Avenue'),
    (3, '789 Pine Road'),
    (4, '101 Maple Lane'),
    (5, '202 Cedar Drive');

INSERT INTO device_types (name) VALUES
    ('Smart Light'),
    ('Smart Promo');

INSERT INTO devices (type_id, house_id, serial_number, status) VALUES
    (1, 1, 'THRM-001', 'Active'),
    (2, 2, 'THRM-002', 'Active'),
    (1, 3, 'THRM-001', 'Active'),
    (1, 4, 'THRM-002', 'Inactive'),
    (2, 5, 'THRM-002', 'Inactive');

INSERT INTO modules (device_id, name) VALUES
    (1, 'Temperature Module'),
    (2, 'Temperature Module'),
    (3, 'Temperature Module'),
    (4, 'Temperature Module'),
    (5, 'Temperature Module');

INSERT INTO telemetry_data (device_id, data, timestamp) VALUES
     (1, '{"temperature": 22.5}', current_timestamp),
     (2, '{"temperature": 17.0}', current_timestamp),
     (3, '{"temperature": 26.9}', current_timestamp),
     (4, '{"temperature": 19.5}', current_timestamp),
     (5, '{"temperature": 20.5}', current_timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd