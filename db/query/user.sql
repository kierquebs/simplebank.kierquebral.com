-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    first_name,
    middle_name,
    last_name,
    email
)values(
    $1,$2,$3,$4,$5,$6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;