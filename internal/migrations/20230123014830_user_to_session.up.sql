alter table sessions
    add "user" text;

alter table sessions
    add constraint sessions_user_fk
        foreign key ("user") references users;
