
-- +goose Up
CREATE TABLE roles
(
    id uuid PRIMARY KEY,
    active boolean,
    name text,
    created_by uuid NOT NULL,
    bound boolean NOT NULL DEFAULT false,
    deleted boolean NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by uuid,
    isolated_entity_id uuid NOT NULL
);

CREATE TABLE role_permissions
(
    role_id uuid NOT NULL REFERENCES roles(id),
    permissions_id uuid NOT NULL REFERENCES permissions(id),
    isolated_entity_id uuid NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS roles;