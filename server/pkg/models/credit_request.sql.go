// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: credit_request.sql

package models

import (
	"context"
	"database/sql"
)

const createCreditRequest = `-- name: CreateCreditRequest :one
INSERT INTO credit_request(
  user_id,promo_used,program,member_id,transaction_time,
  credit_used,reward_should_receive,transaction_status
) VALUES (
  $1, $2, $3, $4, $5, $6,$7,$8
)
RETURNING reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status
`

type CreateCreditRequestParams struct {
	UserID              int32                 `json:"user_id"`
	PromoUsed           sql.NullInt32         `json:"promo_used"`
	Program             int32                 `json:"program"`
	MemberID            string                `json:"member_id"`
	TransactionTime     sql.NullTime          `json:"transaction_time"`
	CreditUsed          float64               `json:"credit_used"`
	RewardShouldReceive float64               `json:"reward_should_receive"`
	TransactionStatus   TransactionStatusEnum `json:"transaction_status"`
}

func (q *Queries) CreateCreditRequest(ctx context.Context, arg CreateCreditRequestParams) (CreditRequest, error) {
	row := q.db.QueryRowContext(ctx, createCreditRequest,
		arg.UserID,
		arg.PromoUsed,
		arg.Program,
		arg.MemberID,
		arg.TransactionTime,
		arg.CreditUsed,
		arg.RewardShouldReceive,
		arg.TransactionStatus,
	)
	var i CreditRequest
	err := row.Scan(
		&i.ReferenceNumber,
		&i.UserID,
		&i.Program,
		&i.MemberID,
		&i.TransactionTime,
		&i.CreditUsed,
		&i.RewardShouldReceive,
		&i.PromoUsed,
		&i.TransactionStatus,
	)
	return i, err
}

const deleteCreditRequest = `-- name: DeleteCreditRequest :exec
DELETE FROM credit_request
WHERE reference_number = $1
`

func (q *Queries) DeleteCreditRequest(ctx context.Context, referenceNumber int64) error {
	_, err := q.db.ExecContext(ctx, deleteCreditRequest, referenceNumber)
	return err
}

const getCreditRequestByID = `-- name: GetCreditRequestByID :one
SELECT reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status FROM credit_request
WHERE reference_number = $1 LIMIT 1
`

func (q *Queries) GetCreditRequestByID(ctx context.Context, referenceNumber int64) (CreditRequest, error) {
	row := q.db.QueryRowContext(ctx, getCreditRequestByID, referenceNumber)
	var i CreditRequest
	err := row.Scan(
		&i.ReferenceNumber,
		&i.UserID,
		&i.Program,
		&i.MemberID,
		&i.TransactionTime,
		&i.CreditUsed,
		&i.RewardShouldReceive,
		&i.PromoUsed,
		&i.TransactionStatus,
	)
	return i, err
}

const getCreditRequestByProg = `-- name: GetCreditRequestByProg :many
SELECT reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status FROM credit_request
WHERE program = $1
`

func (q *Queries) GetCreditRequestByProg(ctx context.Context, program int32) ([]CreditRequest, error) {
	rows, err := q.db.QueryContext(ctx, getCreditRequestByProg, program)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreditRequest
	for rows.Next() {
		var i CreditRequest
		if err := rows.Scan(
			&i.ReferenceNumber,
			&i.UserID,
			&i.Program,
			&i.MemberID,
			&i.TransactionTime,
			&i.CreditUsed,
			&i.RewardShouldReceive,
			&i.PromoUsed,
			&i.TransactionStatus,
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

const getCreditRequestByPromo = `-- name: GetCreditRequestByPromo :many
SELECT reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status FROM credit_request
WHERE program = $1 AND promo_used=$2
`

type GetCreditRequestByPromoParams struct {
	Program   int32         `json:"program"`
	PromoUsed sql.NullInt32 `json:"promo_used"`
}

func (q *Queries) GetCreditRequestByPromo(ctx context.Context, arg GetCreditRequestByPromoParams) ([]CreditRequest, error) {
	rows, err := q.db.QueryContext(ctx, getCreditRequestByPromo, arg.Program, arg.PromoUsed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreditRequest
	for rows.Next() {
		var i CreditRequest
		if err := rows.Scan(
			&i.ReferenceNumber,
			&i.UserID,
			&i.Program,
			&i.MemberID,
			&i.TransactionTime,
			&i.CreditUsed,
			&i.RewardShouldReceive,
			&i.PromoUsed,
			&i.TransactionStatus,
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

const getCreditRequestByUser = `-- name: GetCreditRequestByUser :many
SELECT reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status FROM credit_request
WHERE user_id = $1
`

func (q *Queries) GetCreditRequestByUser(ctx context.Context, userID int32) ([]CreditRequest, error) {
	rows, err := q.db.QueryContext(ctx, getCreditRequestByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreditRequest
	for rows.Next() {
		var i CreditRequest
		if err := rows.Scan(
			&i.ReferenceNumber,
			&i.UserID,
			&i.Program,
			&i.MemberID,
			&i.TransactionTime,
			&i.CreditUsed,
			&i.RewardShouldReceive,
			&i.PromoUsed,
			&i.TransactionStatus,
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

const listCreditRequest = `-- name: ListCreditRequest :many
SELECT reference_number, user_id, program, member_id, transaction_time, credit_used, reward_should_receive, promo_used, transaction_status FROM credit_request
ORDER BY program
`

func (q *Queries) ListCreditRequest(ctx context.Context) ([]CreditRequest, error) {
	rows, err := q.db.QueryContext(ctx, listCreditRequest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreditRequest
	for rows.Next() {
		var i CreditRequest
		if err := rows.Scan(
			&i.ReferenceNumber,
			&i.UserID,
			&i.Program,
			&i.MemberID,
			&i.TransactionTime,
			&i.CreditUsed,
			&i.RewardShouldReceive,
			&i.PromoUsed,
			&i.TransactionStatus,
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

const updateCreditRequest = `-- name: UpdateCreditRequest :exec
UPDATE credit_request 
SET user_id = COALESCE($1,user_id),
program = COALESCE($2,program),
member_id = COALESCE($3,member_id),
transaction_time = COALESCE($4,transaction_time),
credit_used = COALESCE($5,credit_used),
reward_should_receive = COALESCE($6,reward_should_receive),
transaction_status = COALESCE($7,transaction_status),
promo_used = COALESCE($8,promo_used)
WHERE reference_number = $9
`

type UpdateCreditRequestParams struct {
	UserID              int32                 `json:"user_id"`
	Program             int32                 `json:"program"`
	MemberID            string                `json:"member_id"`
	TransactionTime     sql.NullTime          `json:"transaction_time"`
	CreditUsed          float64               `json:"credit_used"`
	RewardShouldReceive float64               `json:"reward_should_receive"`
	TransactionStatus   TransactionStatusEnum `json:"transaction_status"`
	PromoUsed           sql.NullInt32         `json:"promo_used"`
	ReferenceNumber     int64                 `json:"reference_number"`
}

func (q *Queries) UpdateCreditRequest(ctx context.Context, arg UpdateCreditRequestParams) error {
	_, err := q.db.ExecContext(ctx, updateCreditRequest,
		arg.UserID,
		arg.Program,
		arg.MemberID,
		arg.TransactionTime,
		arg.CreditUsed,
		arg.RewardShouldReceive,
		arg.TransactionStatus,
		arg.PromoUsed,
		arg.ReferenceNumber,
	)
	return err
}

const updateTransactionStatusByID = `-- name: UpdateTransactionStatusByID :exec
UPDATE credit_request
SET transaction_status = $1
WHERE reference_number = $2
`

type UpdateTransactionStatusByIDParams struct {
	TransactionStatus TransactionStatusEnum `json:"transaction_status"`
	ReferenceNumber   int64                 `json:"reference_number"`
}

func (q *Queries) UpdateTransactionStatusByID(ctx context.Context, arg UpdateTransactionStatusByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateTransactionStatusByID, arg.TransactionStatus, arg.ReferenceNumber)
	return err
}
