
-- +goose Up
CREATE TABLE users
(
    id uuid PRIMARY KEY,
    active boolean,
    created_at timestamp,
    modified_at timestamp,
    role_id uuid  REFERENCES roles(id),
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;