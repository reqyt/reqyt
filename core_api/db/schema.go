package db

const Schema = `
PRAGMA foreign_keys=1;
CREATE TABLE functions
(
    id integer
        primary key autoincrement not null,
    created_at datetime default 0 not null,
    updated_at datetime default 0 not null,
    deleted_at datetime default null,
    slug varchar(255) not null unique,
    description varchar(255) not null,
    name varchar(255) not null,
    cost unsigned bigint default 0 not null,
    endpoint varchar(255) null,
    wallet_id integer not null,
    foreign key(wallet_id) references wallets(id) DEFERRABLE INITIALLY IMMEDIATE
);
CREATE TABLE users
(
    id integer
        primary key autoincrement not null,
    created_at datetime default 0 not null,
    updated_at datetime default 0 not null,
    deleted_at datetime default null,
    name varchar(255) not null,
    email varchar(255) not null
);
CREATE TABLE wallets
(
    id integer
        primary key autoincrement not null,
    created_at datetime default 0 not null,
    updated_at datetime default 0 not null,
    deleted_at datetime default null,
    key varchar(255) not null,
    credits unsigned bigint default 0 not null
);
CREATE TABLE function_calls
(
    id integer not null
        primary key autoincrement,
    created_at datetime default 0 not null,
    updated_at datetime default 0 not null,
    deleted_at datetime default null,
    function_id integer not null,
    wallet_id integer not null,
    credits unsigned bigint default 0 not null,
    foreign key(function_id) references functions(id) DEFERRABLE INITIALLY IMMEDIATE,
    foreign key(wallet_id) references wallets(id) DEFERRABLE INITIALLY IMMEDIATE
);
CREATE TABLE rel_user_wallet
(
    user_id integer not null,
    wallet_id integer not null,
    primary key (user_id, wallet_id),
    foreign key(user_id) references users(id) deferrable initially immediate,
    foreign key(wallet_id) references wallets(id) deferrable initially immediate
);
`

const Data = `
INSERT INTO wallets(id, key) VALUES (1, "system");
INSERT INTO wallets(id, key) VALUES (2, "public");
INSERT INTO wallets(id, key) VALUES (99, "placeholder");
`