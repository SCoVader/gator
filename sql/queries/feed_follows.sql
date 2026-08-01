-- name: CreateFeedFollow :one
with create_feed_follow as (
insert into feeds_follows(id, created_at, updated_at, user_id, feed_id)
values(
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
) returning *)
select 
feeds.name as feed_name
users.name as user_name
from create_feed_follow
inner join feeds
on feed_follows.feed_id = feeds.id
inner join users
on feed_follows.user_id = users.id
