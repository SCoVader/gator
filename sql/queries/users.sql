-- name: CreateUser :one
insert into "users"(id, created_at, updated_at, name)
values (
    gen_random_uuid(),
    now(),
    now(),
    $1
)
returning *;

-- name: GetUser :one
select * from users
where name = $1; 

-- name: ResetUsers :exec
delete from users;

-- name: GetUsers :many
select name from users;
