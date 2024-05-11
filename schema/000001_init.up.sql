CREATE TABLE users
(
    id                  serial       PRIMARY KEY,
    password_hash       varchar(255) not null,
    email	            varchar(255) not null unique,
    token	            varchar(255)
);

CREATE TABLE users_profile
(
    id                  serial       PRIMARY KEY,
    name                varchar(255) not null,
    surname             varchar(255) not null,
    email varchar(255) not null unique,
    photo               varchar(255) not null,
    city	            varchar(255),
    birthday            date         default NOW()
);

CREATE TABLE chat_lists
(
    id          serial       PRIMARY KEY,
    title       varchar(255) not null
    -- photo 
    -- listchat references chat_items on delete cascade not null 
    -- description varchar(255)
);

CREATE TABLE find_users
(
    id          serial PRIMARY KEY,
    user_id     int references users (id) on delete cascade not null,
    count       varchar(255)    not null,
    start_day   date   not null,
    end_day     date   not null,
    start_time  time   not null,
    end_time    time   not null,
    del         boolean not null default false
);


CREATE TABLE users_hobby
(
    id          serial       PRIMARY KEY,
    description varchar(255) not null
);

CREATE TABLE users_profile_lists
(
    id           serial                                              PRIMARY KEY,
    user_id      int references users (id)         on delete cascade not null,
    profile_id   int references users_profile (id) on delete cascade not null
);

CREATE TABLE users_hobby_lists
(
    id           serial                                              PRIMARY KEY,
    prof_id      int references users_profile (id)         on delete cascade not null,
    userhobby_id int references users_hobby (id)    on delete cascade not null
);

CREATE TABLE users_chat_lists
(
    id           serial                                              PRIMARY KEY,
    user_id      int references users (id)         on delete cascade not null,
    chatlists_id int references chat_lists (id)    on delete cascade not null,
    chatName     varchar(255) not null
);

CREATE TABLE chat_items
(
    id          serial                                           PRIMARY KEY,
    username    varchar(255)                                     not null,
    description                                                  varchar(10000),
    chatlist_id int references chat_lists (id) on delete cascade not null,
    user_id     int references users (id)      on delete cascade not null
);