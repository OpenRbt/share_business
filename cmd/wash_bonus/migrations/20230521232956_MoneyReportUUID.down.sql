alter table session_money_report
    drop constraint session_money_report_unique;

alter table session_money_report
    drop column uuid;
