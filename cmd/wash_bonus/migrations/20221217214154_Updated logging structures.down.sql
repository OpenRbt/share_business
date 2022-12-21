alter table balance_events
    add operation_kind integer not null;


alter table balance_events
    add wash_server uuid not null;


alter table balance_events
    add session text not null;


alter table balance_events
    add status boolean not null;


alter table balance_events
    add error_msg text;


drop table session_event;

