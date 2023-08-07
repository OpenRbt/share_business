alter table sessions
    add column created_at timestamp default now() not null,
    add column updated_at timestamp default now();


