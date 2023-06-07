-- Dropping new constraints
alter table wash_servers
    drop constraint wash_servers_users_id_fk;
alter table users
    drop constraint user_id_pk;

-- Resetting users table to id field

alter table users
    rename column id to identity_uid;
alter table users
    add column id uuid not null default gen_random_uuid();
alter table users
    add constraint user_id_pk
        primary key (id);

-- Updating wash_server relations
alter table wash_servers
    add column owner uuid;

update wash_servers
    set owner = (select "u".id from users "u" where "u".identity_uid = wash_servers.created_by limit 1);

alter table wash_servers
    alter column owner set not null;

alter table wash_servers
    drop column created_by;

alter table wash_servers
    add constraint wash_servers_owner_fk
        foreign key (owner) references users(id);