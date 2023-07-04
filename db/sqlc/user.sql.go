// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    first_name,
    middle_name,
    last_name,
    email
)values(
    $1,$2,$3,$4,$5,$6
) RETURNING username, hashed_password, first_name, middle_name, last_name, email, password_changed_at, created_at, is_email_verified
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, first_name, middle_name, last_name, email, password_changed_at, created_at, is_email_verified FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE($1, hashed_password),
    password_changed_at = COALESCE($2, password_changed_at),
    first_name = COALESCE($3, first_name),
    middle_name = COALESCE($4, middle_name),
    last_name = COALESCE($5, last_name),
    email = COALESCE($6, email),
    is_email_verified = COALESCE($7, is_email_verified)
WHERE
    username = $8
RETURNING username, hashed_password, first_name, middle_name, last_name, email, password_changed_at, created_at, is_email_verified
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	FirstName         sql.NullString `json:"first_name"`
	MiddleName        sql.NullString `json:"middle_name"`
	LastName          sql.NullString `json:"last_name"`
	Email             sql.NullString `json:"email"`
	IsEmailVerified   sql.NullBool   `json:"is_email_verified"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Email,
		arg.IsEmailVerified,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}
