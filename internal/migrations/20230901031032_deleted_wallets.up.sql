ALTER TABLE wallets
    ADD COLUMN deleted boolean NOT NULL DEFAULT false;

UPDATE wallets SET deleted = true WHERE organization_id IN (
    SELECT id
    FROM organizations
    WHERE deleted
);