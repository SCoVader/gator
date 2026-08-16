-- +goose Up
create table "posts"(
    "id" uuid not null primary key,
    "created_at" timestamp not null,
    "updated_at" timestamp not null,
    "title" text,
    "url" text not null,
    "description" text,
    "published_at" timestamp,
    "feed_id" uuid not null references feeds("id")
);

-- +goose Down
drop table if exists "posts";
