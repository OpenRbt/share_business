create table session_money_report
(
    id           SERIAL PRIMARY KEY,
    station_id   INT NOT NULL,
    banknotes    INT default 0 not null,
    cars_total   INT default 0 not null,
    coins        INT default 0 not null,
    electronical INT default 0 not null,
    service      INT default 0 not null,
    bonuses      INT default 0 not null,
    ctime        timestamp default now(),
    processed    BOOLEAN default false not null,
    session_id uuid constraint money_report_session_fk references sessions
);