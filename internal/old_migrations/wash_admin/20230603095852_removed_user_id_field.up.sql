-- Dropping old constraints
alter table wash_servers
    drop constraint wash_servers_owner_fk;
alter table users
    drop constraint user_id_pk;


-- Switching from owner field to created_by
alter table wash_servers
    add created_by text;

update wash_servers
    set created_by = (select "u".identity_uid from users "u" where u.id = wash_servers.owner limit 1);

alter table wash_servers
    alter column created_by set not null;


alter table wash_servers
    drop column owner;

-- Updating users table and creating new constraints

alter table users
    drop column id;

alter table users
    rename column identity_uid to id;

alter table users
    add constraint user_id_pk
        primary key (id);

alter table wash_servers
    add constraint wash_servers_users_id_fk
        foreign key (created_by) references users(id);
