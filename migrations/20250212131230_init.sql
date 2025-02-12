-- +goose Up
create table chats (
    id SERIAL PRIMARY KEY,
    usernames varchar(20)[]
);

create table logs (
    id SERIAL PRIMARY KEY,
    chat_id integer,
    action varchar(255),
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE,
    created_at timestamp not null default now()

);

-- +goose Down
drop table chats;
drop table logs;