alter  table session_money_report
    add column uuid uuid not null default gen_random_uuid();

alter table  session_money_report
    add constraint session_money_report_unique UNIQUE(session_id,station_id,uuid);