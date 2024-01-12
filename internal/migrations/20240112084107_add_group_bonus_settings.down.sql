ALTER TABLE organizations 
    ADD COLUMN processing_delay_interval_temp INTERVAL;

UPDATE organizations
    SET processing_delay_interval_temp = (processing_delay || ' minutes')::interval;

ALTER TABLE organizations 
    DROP COLUMN processing_delay;

ALTER TABLE organizations 
    RENAME COLUMN processing_delay_interval_temp TO processing_delay;

ALTER TABLE organizations 
    ALTER COLUMN processing_delay DROP DEFAULT;

ALTER TABLE server_groups 
    DROP COLUMN IF EXISTS processing_delay;

ALTER TABLE server_groups 
    DROP COLUMN IF EXISTS bonus_percentage;
