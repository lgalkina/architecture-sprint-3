-- +goose Up
-- +goose StatementBegin
INSERT INTO telemetries (device_id, temperature, timestamp) VALUES
     (1, 22.5, current_timestamp),
     (2, 17.0, current_timestamp),
     (3, 26.9, current_timestamp),
     (4, 19.5, current_timestamp),
     (5, 20.5, current_timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd