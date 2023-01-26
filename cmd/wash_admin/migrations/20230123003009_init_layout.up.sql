create table users
(
    id           uuid    default gen_random_uuid() not null
        constraint user_id_pk
            primary key,
    identity_uid text                              not null,
    deleted      boolean default false             not null
);

create table wash_servers
(
    id                  uuid    default gen_random_uuid() not null
        constraint wash_servers_id_pk
            primary key,
    owner               uuid                              not null
        constraint wash_servers_owner_fk
            references users,
    title               text                              not null,
    description         text                              not null,
    deleted             boolean default false             not null
);

