
-- +goose Up
ALTER TABLE IF EXISTS wash_servers ADD COLUMN description text;
ALTER TABLE IF EXISTS wash_servers ADD COLUMN owner_id uuid REFERENCES users (id);
ALTER TABLE IF EXISTS wash_servers ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE IF EXISTS wash_servers RENAME COLUMN key TO service_key;
ALTER TABLE IF EXISTS wash_servers DROP COLUMN isolated_entity_id;
ALTER TABLE IF EXISTS wash_servers DROP COLUMN bound;
ALTER TABLE IF EXISTS wash_servers DROP COLUMN created_by;

-- +goose Down
ALTER TABLE IF EXISTS wash_servers DROP COLUMN description;
ALTER TABLE IF EXISTS wash_servers DROP COLUMN owner_id;
ALTER TABLE IF EXISTS wash_servers ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS wash_servers RENAME COLUMN service_key TO key;
ALTER TABLE IF EXISTS wash_servers ADD COLUMN isolated_entity_id uuid NOT NULL;
ALTER TABLE IF EXISTS wash_servers ADD COLUMN bound boolean NOT NULL DEFAULT false;
ALTER TABLE IF EXISTS wash_servers ADD COLUMN created_by uuid NOT NULL;