ALTER TABLE organizations 
    ADD COLUMN processing_delay_minutes_temp INT;

UPDATE organizations
    SET processing_delay_minutes_temp = EXTRACT(epoch FROM processing_delay) / 60;

ALTER TABLE organizations 
    DROP COLUMN processing_delay;

ALTER TABLE organizations 
    RENAME COLUMN processing_delay_minutes_temp TO processing_delay;

ALTER TABLE organizations 
    ALTER COLUMN processing_delay SET DEFAULT 60;

ALTER TABLE organizations 
    ALTER COLUMN processing_delay SET NOT NULL;

ALTER TABLE server_groups
    ADD COLUMN processing_delay INT,
    ADD COLUMN bonus_percentage INT;