alter table sessions
    drop constraint sessions_user_fk;

alter table sessions
    drop column "user";

