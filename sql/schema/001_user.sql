-- +goose Up
create table "users"(
    "id" uuid not null primary key,
    "created_at" timestamp not null,
    "updated_at" timestamp not null,
    "name" text unique  not null
);

-- +goose Down
drop table if exists "users";
