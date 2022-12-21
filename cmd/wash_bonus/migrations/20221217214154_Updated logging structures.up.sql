alter table balance_events
    drop column operation_kind;

alter table balance_events
    drop column wash_server;

alter table balance_events
    drop column session;

alter table balance_events
    drop column status;

alter table balance_events
    drop column error_msg;

create table session_event
(
    id           uuid default gen_random_uuid() not null
        constraint session_event_pk
            primary key,
    "sessionID"  uuid                           not null,
    "washServer" uuid
        constraint session_event__wash_server_fk
            references wash_servers,
    "postID"     numeric,
    "user"       uuid
        constraint session_event__user_fk
            references users,
    event        text                           not null,
    changes      text
);
