alter table wash_servers
    add column service_key text,
    add column created_by  text
        constraint wash_servers_users_id_fk references users;

alter table users
    add column role text default 'user'::text;