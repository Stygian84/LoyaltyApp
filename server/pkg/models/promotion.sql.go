// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: promotion.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const createPromotion = `-- name: CreatePromotion :one
INSERT INTO promotion(
  program,promo_type,start_date,end_date,
  earn_rate_type,constant, card_tier, loyalty_membership
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, program, promo_type, start_date, end_date, earn_rate_type, constant, card_tier, loyalty_membership
`

type CreatePromotionParams struct {
	Program           int32            `json:"program"`
	PromoType         PromoTypeEnum    `json:"promo_type"`
	StartDate         time.Time        `json:"start_date"`
	EndDate           time.Time        `json:"end_date"`
	EarnRateType      EarnRateTypeEnum `json:"earn_rate_type"`
	Constant          float64          `json:"constant"`
	CardTier          int32            `json:"card_tier"`
	LoyaltyMembership sql.NullInt32    `json:"loyalty_membership"`
}

func (q *Queries) CreatePromotion(ctx context.Context, arg CreatePromotionParams) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, createPromotion,
		arg.Program,
		arg.PromoType,
		arg.StartDate,
		arg.EndDate,
		arg.EarnRateType,
		arg.Constant,
		arg.CardTier,
		arg.LoyaltyMembership,
	)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.Program,
		&i.PromoType,
		&i.StartDate,
		&i.EndDate,
		&i.EarnRateType,
		&i.Constant,
		&i.CardTier,
		&i.LoyaltyMembership,
	)
	return i, err
}

const deletePromotion = `-- name: DeletePromotion :exec
DELETE FROM promotion
WHERE id = $1
`

func (q *Queries) DeletePromotion(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePromotion, id)
	return err
}

const getPromotionByDateRange = `-- name: GetPromotionByDateRange :many
SELECT id, program, promo_type, start_date, end_date, earn_rate_type, constant, card_tier, loyalty_membership FROM promotion
WHERE $1 >= start_date AND $1 <= end_date AND program=$2
`

type GetPromotionByDateRangeParams struct {
	Column1 interface{} `json:"column_1"`
	Program int32       `json:"program"`
}

func (q *Queries) GetPromotionByDateRange(ctx context.Context, arg GetPromotionByDateRangeParams) ([]Promotion, error) {
	rows, err := q.db.QueryContext(ctx, getPromotionByDateRange, arg.Column1, arg.Program)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Promotion
	for rows.Next() {
		var i Promotion
		if err := rows.Scan(
			&i.ID,
			&i.Program,
			&i.PromoType,
			&i.StartDate,
			&i.EndDate,
			&i.EarnRateType,
			&i.Constant,
			&i.CardTier,
			&i.LoyaltyMembership,
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

const getPromotionByID = `-- name: GetPromotionByID :one
SELECT id, program, promo_type, start_date, end_date, earn_rate_type, constant, card_tier, loyalty_membership FROM promotion
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPromotionByID(ctx context.Context, id int64) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, getPromotionByID, id)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.Program,
		&i.PromoType,
		&i.StartDate,
		&i.EndDate,
		&i.EarnRateType,
		&i.Constant,
		&i.CardTier,
		&i.LoyaltyMembership,
	)
	return i, err
}

const getPromotionByProg = `-- name: GetPromotionByProg :many
SELECT id, program, promo_type, start_date, end_date, earn_rate_type, constant, card_tier, loyalty_membership FROM promotion
WHERE program = $1
`

func (q *Queries) GetPromotionByProg(ctx context.Context, program int32) ([]Promotion, error) {
	rows, err := q.db.QueryContext(ctx, getPromotionByProg, program)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Promotion
	for rows.Next() {
		var i Promotion
		if err := rows.Scan(
			&i.ID,
			&i.Program,
			&i.PromoType,
			&i.StartDate,
			&i.EndDate,
			&i.EarnRateType,
			&i.Constant,
			&i.CardTier,
			&i.LoyaltyMembership,
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

const listPromotion = `-- name: ListPromotion :many
SELECT promotion.id, program, promo_type, start_date, end_date, earn_rate_type, constant, card_tier, loyalty_membership, loyalty_program.id, name, currency_name, processing_time, description, enrollment_link, terms_condition_link, format_regex, partner_code, initial_earn_rate FROM promotion LEFT JOIN loyalty_program ON promotion.program= loyalty_program.id
`

type ListPromotionRow struct {
	ID                 int64            `json:"id"`
	Program            int32            `json:"program"`
	PromoType          PromoTypeEnum    `json:"promo_type"`
	StartDate          time.Time        `json:"start_date"`
	EndDate            time.Time        `json:"end_date"`
	EarnRateType       EarnRateTypeEnum `json:"earn_rate_type"`
	Constant           float64          `json:"constant"`
	CardTier           int32            `json:"card_tier"`
	LoyaltyMembership  sql.NullInt32    `json:"loyalty_membership"`
	ID_2               sql.NullInt64    `json:"id_2"`
	Name               sql.NullString   `json:"name"`
	CurrencyName       sql.NullString   `json:"currency_name"`
	ProcessingTime     sql.NullString   `json:"processing_time"`
	Description        sql.NullString   `json:"description"`
	EnrollmentLink     sql.NullString   `json:"enrollment_link"`
	TermsConditionLink sql.NullString   `json:"terms_condition_link"`
	FormatRegex        sql.NullString   `json:"format_regex"`
	PartnerCode        sql.NullString   `json:"partner_code"`
	InitialEarnRate    sql.NullFloat64  `json:"initial_earn_rate"`
}

func (q *Queries) ListPromotion(ctx context.Context) ([]ListPromotionRow, error) {
	rows, err := q.db.QueryContext(ctx, listPromotion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPromotionRow
	for rows.Next() {
		var i ListPromotionRow
		if err := rows.Scan(
			&i.ID,
			&i.Program,
			&i.PromoType,
			&i.StartDate,
			&i.EndDate,
			&i.EarnRateType,
			&i.Constant,
			&i.CardTier,
			&i.LoyaltyMembership,
			&i.ID_2,
			&i.Name,
			&i.CurrencyName,
			&i.ProcessingTime,
			&i.Description,
			&i.EnrollmentLink,
			&i.TermsConditionLink,
			&i.FormatRegex,
			&i.PartnerCode,
			&i.InitialEarnRate,
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

const updatePromotion = `-- name: UpdatePromotion :exec
UPDATE promotion 
SET program = COALESCE($1,program),
promo_type = COALESCE($2,promo_type),
start_date = COALESCE($3,start_date),
end_date = COALESCE($4,end_date),
earn_rate_type = COALESCE($5,earn_rate_type),
constant = COALESCE($6,constant),
card_tier = COALESCE($7,card_tier),
loyalty_membership = COALESCE($8,loyalty_membership)
WHERE id = $9
`

type UpdatePromotionParams struct {
	Program           int32            `json:"program"`
	PromoType         PromoTypeEnum    `json:"promo_type"`
	StartDate         time.Time        `json:"start_date"`
	EndDate           time.Time        `json:"end_date"`
	EarnRateType      EarnRateTypeEnum `json:"earn_rate_type"`
	Constant          float64          `json:"constant"`
	CardTier          int32            `json:"card_tier"`
	LoyaltyMembership sql.NullInt32    `json:"loyalty_membership"`
	ID                int64            `json:"id"`
}

func (q *Queries) UpdatePromotion(ctx context.Context, arg UpdatePromotionParams) error {
	_, err := q.db.ExecContext(ctx, updatePromotion,
		arg.Program,
		arg.PromoType,
		arg.StartDate,
		arg.EndDate,
		arg.EarnRateType,
		arg.Constant,
		arg.CardTier,
		arg.LoyaltyMembership,
		arg.ID,
	)
	return err
}
