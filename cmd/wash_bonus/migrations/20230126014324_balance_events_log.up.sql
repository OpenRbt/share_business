create table balance_events
(
    id         uuid default gen_random_uuid() not null
        constraint balance_events_pk
            primary key,
    "user"     uuid                           not null
        constraint balance_events_users_fk
            references users (id),
    old_amount numeric(10, 2)                 not null,
    new_amount numeric(10, 2),
    date       timestamp                      not null
);

create table sessions_balance_events
(
    id         uuid default gen_random_uuid() not null
        constraint sessions_balance_events_pk
            primary key,
    "session"     uuid                           not null
        constraint sessions_balance_events_sessions_fk
            references sessions (id),
    old_amount numeric(10, 2)                 not null,
    new_amount numeric(10, 2),
    date       timestamp                      not null
);

