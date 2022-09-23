
-- +goose Up
CREATE TABLE bonus_balance
(
    id uuid PRIMARY KEY,
    user_id REFERENCES users(id),
    balance float,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
);


-- +goose Down
DROP TABLE IF EXISTS bonus_balance;