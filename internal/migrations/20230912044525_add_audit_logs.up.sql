CREATE TYPE resource_type AS ENUM ('admin_applications', 'admin_users', 'organizations', 'server_groups', 'wash_servers');

CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    resource resource_type NOT NULL,
    entity_id TEXT NOT NULL,
    action TEXT NOT NULL,
    user_performing_action TEXT NOT NULL,
    performed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data JSONB
);
