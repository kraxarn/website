create table users
(
    id       serial  not null primary key,
    username text    not null unique,
    password bytea   not null,
    flags    integer not null
);

create table texts
(
    id        serial      not null primary key,
    key       text        not null,
    value     text        not null,
    editor    serial      not null references users (id),
    timestamp timestamptz not null
);