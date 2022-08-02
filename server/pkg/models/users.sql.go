// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package models

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
  full_name,credit_balance,email,contact,
  password,user_name,card_tier
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at
`

type CreateUserParams struct {
	FullName      sql.NullString `json:"full_name"`
	CreditBalance float64        `json:"credit_balance"`
	Email         string         `json:"email"`
	Contact       sql.NullInt32  `json:"contact"`
	Password      string         `json:"password"`
	UserName      string         `json:"user_name"`
	CardTier      sql.NullInt32  `json:"card_tier"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FullName,
		arg.CreditBalance,
		arg.Email,
		arg.Contact,
		arg.Password,
		arg.UserName,
		arg.CardTier,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreditBalance,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserName,
		&i.CardTier,
		&i.CreatedAt,
	)
	return i, err
}

const decreBalance = `-- name: DecreBalance :exec
UPDATE users 
SET credit_balance = credit_balance-$1 
WHERE id=$2 and credit_balance>=$1
`

type DecreBalanceParams struct {
	CreditBalance float64 `json:"credit_balance"`
	ID            int64   `json:"id"`
}

func (q *Queries) DecreBalance(ctx context.Context, arg DecreBalanceParams) error {
	_, err := q.db.ExecContext(ctx, decreBalance, arg.CreditBalance, arg.ID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreditBalance,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserName,
		&i.CardTier,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUserCardTier = `-- name: GetUserByUserCardTier :many
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
WHERE card_tier = $1
`

func (q *Queries) GetUserByUserCardTier(ctx context.Context, cardTier sql.NullInt32) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUserByUserCardTier, cardTier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreditBalance,
			&i.Email,
			&i.Contact,
			&i.Password,
			&i.UserName,
			&i.CardTier,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUserContact = `-- name: GetUserByUserContact :many
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
WHERE contact = $1
`

func (q *Queries) GetUserByUserContact(ctx context.Context, contact sql.NullInt32) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUserByUserContact, contact)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreditBalance,
			&i.Email,
			&i.Contact,
			&i.Password,
			&i.UserName,
			&i.CardTier,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUserEmail = `-- name: GetUserByUserEmail :many
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByUserEmail(ctx context.Context, email string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUserByUserEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreditBalance,
			&i.Email,
			&i.Contact,
			&i.Password,
			&i.UserName,
			&i.CardTier,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetUserByUserName(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, userName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreditBalance,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserName,
		&i.CardTier,
		&i.CreatedAt,
	)
	return i, err
}

const incrBalance = `-- name: IncrBalance :exec
UPDATE users 
SET credit_balance = credit_balance+$1 
WHERE id=$2
`

type IncrBalanceParams struct {
	CreditBalance float64 `json:"credit_balance"`
	ID            int64   `json:"id"`
}

func (q *Queries) IncrBalance(ctx context.Context, arg IncrBalanceParams) error {
	_, err := q.db.ExecContext(ctx, incrBalance, arg.CreditBalance, arg.ID)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, full_name, credit_balance, email, contact, password, user_name, card_tier, created_at FROM users
ORDER BY user_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreditBalance,
			&i.Email,
			&i.Contact,
			&i.Password,
			&i.UserName,
			&i.CardTier,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
SET full_name = COALESCE($1,full_name),
credit_balance = COALESCE($2,credit_balance),
email = COALESCE($3,email),
contact = COALESCE($4,contact),
password = COALESCE($5,password),
user_name = COALESCE($6,user_name),
card_tier = COALESCE($7,card_tier)
WHERE id = $8
`

type UpdateUserParams struct {
	FullName      sql.NullString `json:"full_name"`
	CreditBalance float64        `json:"credit_balance"`
	Email         string         `json:"email"`
	Contact       sql.NullInt32  `json:"contact"`
	Password      string         `json:"password"`
	UserName      string         `json:"user_name"`
	CardTier      sql.NullInt32  `json:"card_tier"`
	ID            int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FullName,
		arg.CreditBalance,
		arg.Email,
		arg.Contact,
		arg.Password,
		arg.UserName,
		arg.CardTier,
		arg.ID,
	)
	return err
}
