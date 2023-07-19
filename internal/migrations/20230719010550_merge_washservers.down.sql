alter table users
    drop column role;

alter table wash_servers
    drop column service_key,
    drop column created_by;