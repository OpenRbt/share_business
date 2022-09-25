
-- +goose Up
CREATE TABLE users
(
    id uuid PRIMARY KEY default gen_random_uuid(),
    active boolean,
    identity_id text,
    created_at timestamp,
    modified_at timestamp,
    modified_by uuid,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid
);


-- +goose Down
DROP TABLE IF EXISTS users;