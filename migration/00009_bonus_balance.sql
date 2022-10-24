
-- +goose Up
CREATE TABLE bonus_balance
(
    id uuid PRIMARY KEY,
    user_id uuid REFERENCES users(id),
    balance numeric(12, 2),
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid
);


-- +goose Down
DROP TABLE IF EXISTS bonus_balance;