-- +goose Up
-- +goose StatementBegin
ALTER TABLE cartridges
ALTER COLUMN barcode_number TYPE INTEGER USING (barcode_number::INTEGER);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cartridges
ALTER COLUMN barcode_number TYPE VARCHAR(55) USING (barcode_number::VARCHAR);
-- +goose StatementEnd