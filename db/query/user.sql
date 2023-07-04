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

-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    first_name = COALESCE(sqlc.narg(first_name), first_name),
    middle_name = COALESCE(sqlc.narg(middle_name), middle_name),
    last_name = COALESCE(sqlc.narg(last_name), last_name),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
    username = sqlc.arg(username)
RETURNING *;