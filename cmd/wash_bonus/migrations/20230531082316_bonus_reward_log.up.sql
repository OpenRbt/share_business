create table bonus_reward_log
(
    id         bigserial not null
        constraint bonus_reward_log_pk
            primary key,
    amount    numeric(10,2),
    session_id uuid
        constraint bonus_reward_log_sessions_null_fk
            references sessions (id),
    uuid uuid not null
);

alter table  bonus_reward_log
    add constraint bonus_reward_log_unique UNIQUE(session_id,uuid);