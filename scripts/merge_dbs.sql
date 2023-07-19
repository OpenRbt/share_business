CREATE EXTENSION IF NOT EXISTS dblink;

UPDATE wash_servers
SET 
    service_key = sub.service_key,
    created_by = sub.created_by
FROM (
    SELECT *
    FROM dblink('dbname=sample host=localhost port=5432 user=sample password=sample',
      'SELECT id, service_key, created_by FROM wash_servers'
   )
    AS remote_data(id uuid, service_key text, created_by text)
) AS sub
WHERE wash_servers.id = sub.id;

INSERT INTO users (id, deleted, role)
SELECT *
FROM (
    SELECT *
    FROM dblink('dbname=sample host=localhost port=5432 user=sample password=sample',
      'SELECT id, deleted, role FROM users')
    AS remote_data(id text, deleted boolean, role text)
) AS sub
WHERE NOT EXISTS (
    SELECT 1
    FROM users
    WHERE users.id = sub.id
);

UPDATE users
SET 
    role = sub.role
FROM (
    SELECT *
    FROM dblink('dbname=sample host=localhost port=5432 user=sample password=sample',
      'SELECT id, role FROM users'
   )
    AS remote_data(id text, role text)
) AS sub
WHERE users.id = sub.id;