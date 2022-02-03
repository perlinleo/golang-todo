SET SYNCHRONOUS_COMMIT = 'off';
create extension if not exists citext;

DROP TABLE IF EXISTS users CASCADE;
CREATE UNLOGGED TABLE IF NOT EXISTS users
(
    user_id       serial not null primary key,
    user_nickname citext COLLATE "POSIX" not null unique,
    user_email citext  COLLATE "POSIX" not null unique,
    user_fullname varchar(100)   not null
    CHECK (user_nickname <> '' or user_email <> '')
);

DROP TABLE IF EXISTS tasks;
CREATE UNLOGGED TABLE IF NOT EXISTS forums
(
    task_id       serial not null primary key,
    user_id     serial   REFERENCES users(user_id) not null,
    task_name citext NOT NULL,
    task_done boolean NOT NULL default false,
    task_description citext NULL
);


