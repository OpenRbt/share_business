CREATE TABLE organization_settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    organization_id uuid NOT NULL UNIQUE REFERENCES organizations(id),
    processing_delay INTERVAL NOT NULL DEFAULT '60 minutes',
    bonus_percentage INT NOT NULL DEFAULT 5
);

INSERT INTO organization_settings (organization_id, processing_delay, bonus_percentage)
    SELECT id, processing_delay, bonus_percentage
    FROM organizations;

ALTER TABLE organizations
    DROP COLUMN processing_delay,
    DROP COLUMN bonus_percentage,
    DROP COLUMN display_name;

DROP TABLE admin_applications;

ALTER TABLE wash_servers
    ADD CONSTRAINT wash_servers_users_id_fk FOREIGN KEY (created_by) REFERENCES users(id);

ALTER TABLE wash_servers
    DROP CONSTRAINT wash_servers_admin_users_id_fk;

DROP TABLE admin_users;

DROP TYPE ADMIN_ROLE_ENUM;
DROP TYPE ADMIN_APPLICATION_STATUS;

ALTER TABLE users
    ADD COLUMN role TEXT DEFAULT 'user'::TEXT;

CREATE TABLE organization_managers
(
    id              uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    user_id         TEXT NOT NULL REFERENCES users(id),
    organization_id uuid NOT NULL REFERENCES organizations(id),
    UNIQUE (user_id, organization_id)
);