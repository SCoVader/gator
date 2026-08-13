-- name: CreateFeedFollow :one
with create_feed_follow as (
insert into feed_follows(id, created_at, updated_at, user_id, feed_id)
values(
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
) returning *)
select
create_feed_follow.*,
feeds.name as feed_name,
users.name as user_name
from create_feed_follow
inner join feeds
on create_feed_follow.feed_id = feeds.id
inner join users
on create_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
with get_feed_follows as (
    select * from feed_follows
    where feed_follows.user_id = $1
)
select get_feed_follows.*,
feeds.name as feed_name,
users.name as user_name
from get_feed_follows
inner join feeds
on get_feed_follows.feed_id = feeds.id
inner join users
on get_feed_follows.user_id = users.id;

-- name: Unfollow :exec
delete from feed_follows
where user_id = $1 and feed_id = $2;
