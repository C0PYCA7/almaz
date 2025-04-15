-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS cartridges (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    status VARCHAR(55),
    received_from VARCHAR(100),
    received_from_subdivision_date TIMESTAMP,
    send_to_refilling_date TIMESTAMP,
    received_from_refilling_date TIMESTAMP,
    send_to VARCHAR(100),
    send_to_subdivision_date TIMESTAMP,
    barcode_number VARCHAR(55) UNIQUE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS cartridges;
-- +goose StatementEnd
