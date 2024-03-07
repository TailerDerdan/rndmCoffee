CREATE TABLE IF NOT EXISTS chat_items (
    id          serial       PRIMARY KEY,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);
CREATE TABLE IF NOT EXISTS lists_items (
    id      serial                                           PRIMARY KEY,
    item_id int references chat_items (id) on delete cascade not null,
    list_id int references chat_lists (id) on delete cascade not null
);
CREATE TABLE IF NOT EXISTS users_lists(
    id      serial                                           PRIMARY KEY,
    user_id int references users (id) on delete cascade      not null,
    list_id int references chat_lists (id) on delete cascade not null
);
CREATE TABLE IF NOT EXISTS chat_lists (
    id          serial       PRIMARY KEY,
    title       varchar(255) not null,
    description varchar(255)
);
CREATE TABLE IF NOT EXISTS users (
    id            serial       PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);
