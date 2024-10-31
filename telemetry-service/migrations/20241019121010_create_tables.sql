-- +goose Up
-- +goose StatementBegin
CREATE TABLE telemetries (
    id SERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    temperature DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd