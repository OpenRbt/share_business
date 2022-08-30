
-- +goose Up
CREATE TABLE wash_sessions
(
    id uuid PRIMARY KEY,
    active boolean,
    closing_at timestamp,
    created_at timestamp,
    expiration_at timestamp,
    update_at timestamp,
    user_id uuid  REFERENCES tokens(id),
    wash_server_id uuid  REFERENCES wash_servers(id),
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS wash_sessions;