
-- +goose Up
CREATE TABLE users
(
    id uuid PRIMARY KEY,
    active boolean,
    created_at timestamp,
    firebase_id text,
    modified_at timestamp,
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;