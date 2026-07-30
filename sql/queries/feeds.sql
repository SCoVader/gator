-- name: CreateFeed :one
insert into feeds(id, created_at, updated_at, name, url, user_id)
values(
  gen_random_uuid(),
  now(),
  now(),
  $1,
  $2,
  $3  
) returning *;

-- name: GetFeed :one
select * from feeds
where name = $1;

-- name: ResetFeeds :exec
delete from feeds;

-- name: GetFeeds :many
select * from feeds;

-- name: GetFeedsShort :many
select name, url, (select name from users where id = user_id) as added_by
from feeds;
