create table users
(
    id           text                                     not null
        constraint user_id_pk
            primary key,
    balance      numeric(10, 2) default 0                 not null,
    deleted      boolean        default false             not null
);

create table wash_servers
(
    id          uuid    default gen_random_uuid() not null
        constraint wash_servers_id_pk
            primary key,
    title       text                              not null,
    description text                              not null,
    deleted     boolean default false             not null
);

create table sessions
(
    ID          uuid default gen_random_uuid() not null
        constraint sessions_id_pk
            primary key,
    wash_server uuid                           not null
        constraint session_server_fk
            references wash_servers (id),
    post_id     numeric,
    started     bool default false             not null,
    finished    bool default false             not null
);
