// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: loyalty_membership.sql

package models

import (
	"context"
)

const createLoyaltyMembership = `-- name: CreateLoyaltyMembership :one
INSERT INTO loyalty_membership(
  program,name
) VALUES (
  $1, $2 
)
RETURNING id, program, name
`

type CreateLoyaltyMembershipParams struct {
	Program int32  `json:"program"`
	Name    string `json:"name"`
}

func (q *Queries) CreateLoyaltyMembership(ctx context.Context, arg CreateLoyaltyMembershipParams) (LoyaltyMembership, error) {
	row := q.db.QueryRowContext(ctx, createLoyaltyMembership, arg.Program, arg.Name)
	var i LoyaltyMembership
	err := row.Scan(&i.ID, &i.Program, &i.Name)
	return i, err
}

const deleteLoyaltyMembership = `-- name: DeleteLoyaltyMembership :exec
DELETE FROM loyalty_membership
WHERE id = $1
`

func (q *Queries) DeleteLoyaltyMembership(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLoyaltyMembership, id)
	return err
}

const getLoyaltyMembershipByID = `-- name: GetLoyaltyMembershipByID :one
SELECT id, program, name FROM loyalty_membership
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLoyaltyMembershipByID(ctx context.Context, id int64) (LoyaltyMembership, error) {
	row := q.db.QueryRowContext(ctx, getLoyaltyMembershipByID, id)
	var i LoyaltyMembership
	err := row.Scan(&i.ID, &i.Program, &i.Name)
	return i, err
}

const getLoyaltyMembershipByProgram = `-- name: GetLoyaltyMembershipByProgram :many
SELECT id, program, name FROM loyalty_membership
WHERE program = $1
`

func (q *Queries) GetLoyaltyMembershipByProgram(ctx context.Context, program int32) ([]LoyaltyMembership, error) {
	rows, err := q.db.QueryContext(ctx, getLoyaltyMembershipByProgram, program)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LoyaltyMembership
	for rows.Next() {
		var i LoyaltyMembership
		if err := rows.Scan(&i.ID, &i.Program, &i.Name); err != nil {
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

const listLoyaltyMembership = `-- name: ListLoyaltyMembership :many
SELECT id, program, name FROM loyalty_membership
ORDER BY program
`

func (q *Queries) ListLoyaltyMembership(ctx context.Context) ([]LoyaltyMembership, error) {
	rows, err := q.db.QueryContext(ctx, listLoyaltyMembership)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LoyaltyMembership
	for rows.Next() {
		var i LoyaltyMembership
		if err := rows.Scan(&i.ID, &i.Program, &i.Name); err != nil {
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

const updateLoyaltyMembership = `-- name: UpdateLoyaltyMembership :exec
UPDATE loyalty_membership 
SET name = COALESCE($1,name),
program = COALESCE($3,program)
WHERE id = $2
`

type UpdateLoyaltyMembershipParams struct {
	Name    string `json:"name"`
	ID      int64  `json:"id"`
	Program int32  `json:"program"`
}

func (q *Queries) UpdateLoyaltyMembership(ctx context.Context, arg UpdateLoyaltyMembershipParams) error {
	_, err := q.db.ExecContext(ctx, updateLoyaltyMembership, arg.Name, arg.ID, arg.Program)
	return err
}
