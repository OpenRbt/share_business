CREATE TABLE organization_settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    organization_id uuid NOT NULL UNIQUE REFERENCES organizations(id),
    processing_delay INTERVAL NOT NULL DEFAULT '60 minutes',
    bonus_percentage INT NOT NULL DEFAULT 5
);

INSERT INTO organization_settings (organization_id)
    SELECT id FROM organizations;