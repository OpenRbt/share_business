DROP TABLE organization_managers;

ALTER TABLE users
    DROP COLUMN role;

CREATE TYPE ADMIN_ROLE_ENUM AS ENUM ('system_manager', 'admin');
CREATE TYPE ADMIN_APPLICATION_STATUS AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE admin_users (
    id              TEXT            PRIMARY KEY,
    name            TEXT            NOT NULL,
    email           TEXT            NOT NULL,
    role            ADMIN_ROLE_ENUM NOT NULL DEFAULT 'admin'::ADMIN_ROLE_ENUM,
    organization_id uuid REFERENCES organizations(id),
    deleted         BOOLEAN         NOT NULL DEFAULT false
);

ALTER TABLE wash_servers
    DROP CONSTRAINT wash_servers_users_id_fk;

INSERT INTO admin_users (id, name, email, role)
    SELECT DISTINCT ON (u.id) u.id, COALESCE(u.name, ''), COALESCE(u.email, ''), 'system_manager'::ADMIN_ROLE_ENUM
    FROM wash_servers w
    JOIN users u ON w.created_by = u.id;

ALTER TABLE wash_servers
    ADD CONSTRAINT wash_servers_admin_users_id_fk FOREIGN KEY (created_by) REFERENCES admin_users(id);

CREATE TABLE admin_applications (
    id             uuid                     NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_user_id  TEXT                     NOT NULL UNIQUE,
    name           TEXT                     NOT NULL,
    email          TEXT                     NOT NULL,
    status         ADMIN_APPLICATION_STATUS NOT NULL DEFAULT 'pending'::ADMIN_APPLICATION_STATUS
);

ALTER TABLE organizations
    ADD COLUMN processing_delay INTERVAL NOT NULL DEFAULT '60 minutes',
    ADD COLUMN bonus_percentage INT      NOT NULL DEFAULT 5,
    ADD COLUMN display_name     TEXT;

UPDATE organizations o 
    SET processing_delay = s.processing_delay,
        bonus_percentage = s.bonus_percentage,
        display_name     = o.name
    FROM organization_settings s 
    WHERE o.id = s.organization_id;


WITH NumberedOrganizations AS (
    SELECT 
        id, 
        name, 
        ROW_NUMBER() OVER(PARTITION BY name ORDER BY id) AS row
    FROM organizations
)

UPDATE organizations o
    SET display_name = o.name || ' ' || n.row
    FROM NumberedOrganizations n
    WHERE o.id = n.id AND n.row > 1;

ALTER TABLE organizations
    ADD CONSTRAINT org_display_name_uq UNIQUE (display_name);

ALTER TABLE organizations
    ALTER COLUMN display_name SET NOT NULL;

DROP TABLE organization_settings;