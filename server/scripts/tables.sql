create table accounts (
    id text primary key
    mail_address varchar(100) not null,
    hashed_password varchar(100) not null,
    is_deleted boolean default false
)

create table reserves (
    id int generated always as identity primary key
    mail_address varchar(100) not null,
    exp timestamp with time zone default current_timestamp,
    is_deleted boolean default false
)