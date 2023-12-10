create table organizations
(
    id          uuid default gen_random_uuid() not null primary key,
    name        text not null,
    description text null,
    is_default  boolean default false not null,
    deleted     boolean default false not null,
    cost_per_day numeric(10, 2) default 0 not null
);

do $$
declare
    v_organization_id uuid;
    v_group_id     uuid;
begin
    insert into organizations (name, description, is_default)
    values ('Default Organization', 'Default', true)
    returning id into v_organization_id;

    alter table session_money_report
        add column organization_id uuid references organizations(id);

    update session_money_report
        set organization_id = v_organization_id;

    alter table session_money_report
        alter column organization_id set not null;
    
    create table server_groups
    (
        id              uuid default gen_random_uuid() not null primary key,
        organization_id uuid not null references organizations(id),
        name            text not null,
        description     text null,
        is_default      boolean default false not null,
        deleted         boolean default false not null,
        cost_per_day numeric(10, 2) default 0 not null
    );

    insert into server_groups (organization_id, name, description, is_default)
    values (v_organization_id, 'Default Server Group', 'Default', true)
    returning id into v_group_id;
    
    alter table wash_servers
        add column group_id uuid;

    update wash_servers
        set group_id = v_group_id;

    alter table wash_servers
        alter column group_id set not null;

    alter table wash_servers
        add constraint wash_servers_server_groups_id_fk foreign key (group_id) references server_groups;
    
    create table wallets
    (
        id              uuid default gen_random_uuid() not null primary key,
        user_id         text not null references users(id),
        organization_id uuid not null references organizations(id),
        balance         numeric(10, 2) default 0 not null,
        is_default      boolean default false not null
    );
    
    insert into wallets (user_id, organization_id, balance, is_default)
    select
        id as user_id,
        v_organization_id as organization_id,
        balance,
        true as is_default
    from users;

    alter table users
        add column email text,
        add column first_name text,
        add column second_name text,
        drop column balance;

    create table organization_managers
    (
        id              uuid default gen_random_uuid() not null primary key,
        user_id         text not null references users(id),
        organization_id uuid not null references organizations(id),
        unique (user_id, organization_id)
    );

    alter table balance_events
        add column wallet_id uuid references wallets(id);

    update balance_events as ev
        set wallet_id = (select w.id from wallets as w where w.user_id = ev.user and is_default);

    alter table balance_events
        alter column wallet_id set not null;
end $$;
