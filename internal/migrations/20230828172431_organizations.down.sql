alter table users
    drop column email,
    drop column first_name,
    drop column second_name,
    add column balance numeric(10, 2) not null default 0;

update users as u set u.balance = (select sum(w.balance) from wallets as w where w.user_id = u.id);

drop table wallets;
alter table wash_servers drop column group_id;
alter table session_money_reports drop column organization_id;
alter table balance_events drop column wallet_id;
drop table server_groups;
drop table organizations;