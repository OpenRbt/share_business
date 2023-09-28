ALTER TYPE ADMIN_ROLE_ENUM
    ADD VALUE 'no_access';

ALTER TABLE admin_users
    DROP COLUMN deleted;