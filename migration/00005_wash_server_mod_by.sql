
-- +goose Up
ALTER TABLE IF EXISTS wash_servers ADD COLUMN modified_by uuid REFERENCES users (id);

-- +goose Down
ALTER TABLE IF EXISTS wash_servers DROP COLUMN modified_by;