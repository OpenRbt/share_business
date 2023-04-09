alter table sessions
    add column balance numeric(10, 2) default 0 not null;