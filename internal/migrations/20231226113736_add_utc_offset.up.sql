ALTER TABLE organizations
    ADD COLUMN utc_offset INTEGER NOT NULL DEFAULT 0 CHECK (utc_offset >= -720 AND utc_offset <= 840);

ALTER TABLE server_groups
    ADD COLUMN utc_offset INTEGER CHECK (utc_offset >= -720 AND utc_offset <= 840);
