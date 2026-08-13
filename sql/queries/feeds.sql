-- name: CreateFeed :one
insert into feeds(id, created_at, updated_at, last_fetched_at, name, url, user_id)
values(
  gen_random_uuid(),
  now(),
  now(),
  now(),
  $1,
  $2,
  $3  
) returning *;

-- name: UpdateFetchedStatus :one
update feeds set updated_at = now(), last_fetched_at = now()
where id = $1;

-- name: GetFeed :one
select * from feeds
where name = $1;

-- name: GetFeedByURL :one
select * from feeds
where url = $1;

-- name: ResetFeeds :exec
delete from feeds;

-- name: GetFeeds :many
select * from feeds;

-- name: GetFeedsShort :many
select name, url, (select name from users where id = user_id) as added_by
from feeds;
