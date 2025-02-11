-- +goose Up
create table chats (
    id SERIAL PRIMARY KEY,
    usernames varchar(10)[]
);

-- +goose Down
drop table chats;
