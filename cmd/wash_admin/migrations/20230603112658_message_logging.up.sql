create table rabbit_send_log
(
    id           bigserial          not null
        constraint rabbit_send_log_pk
            primary key,
    message_type text               not null,
    payload      jsonb              not null,
    created_at   timestamp          not null,
    sent         bool default false not null,
    sent_at      timestamp
);

create index rabbit_send_log_sent_index
    on rabbit_send_log (sent);