create table users
(
    id       uuid           default gen_random_uuid() not null
        primary key,
    identity text                                     not null,
    balance  numeric(10, 2) default 0                 not null,
    active   boolean        default true              not null
);

create table wash_admins
(
    id       uuid default gen_random_uuid() not null
        constraint wash_admins_pk
            primary key,
    identity text                           not null
);

create table wash_servers
(
    id          uuid default gen_random_uuid() not null
        constraint wash_servers_pk
            primary key,
    owner       uuid                           not null
        constraint wash_servers_wash_admins_null_fk
            references wash_admins,
    name        text                           not null,
    description text,
    wash_key    text                           not null
);

create table balance_events
(
    id             uuid default gen_random_uuid() not null
        constraint balance_events_pk
            primary key,
    "user"         uuid                           not null
        constraint balance_events_users_null_fk
            references users,
    operation_kind integer                        not null,
    old_amount     numeric(10, 2)                 not null,
    new_amount     numeric(10, 2)                 not null,
    wash_server    uuid                           not null,
    session        text                           not null,
    date           timestamp                      not null
);
