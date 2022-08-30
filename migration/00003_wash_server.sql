
-- +goose Up
CREATE TABLE wash_servers
(
    id uuid PRIMARY KEY,
    created_at timestamp,
    key text,
    last_update_at timestamp,
    modified_at timestamp,
    name text,
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS wash_servers;