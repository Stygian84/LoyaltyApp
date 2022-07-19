// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package models

import (
	"database/sql"
	"fmt"
	"time"
)

type EarnRateTypeEnum string

const (
	EarnRateTypeEnumAdd EarnRateTypeEnum = "add"
	EarnRateTypeEnumMul EarnRateTypeEnum = "mul"
)

func (e *EarnRateTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EarnRateTypeEnum(s)
	case string:
		*e = EarnRateTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for EarnRateTypeEnum: %T", src)
	}
	return nil
}

type PromoTypeEnum string

const (
	PromoTypeEnumOnetime PromoTypeEnum = "onetime"
	PromoTypeEnumOngoing PromoTypeEnum = "ongoing"
)

func (e *PromoTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PromoTypeEnum(s)
	case string:
		*e = PromoTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for PromoTypeEnum: %T", src)
	}
	return nil
}

type TransactionStatusEnum string

const (
	TransactionStatusEnumCreated  TransactionStatusEnum = "created"
	TransactionStatusEnumPending  TransactionStatusEnum = "pending"
	TransactionStatusEnumApproved TransactionStatusEnum = "approved"
	TransactionStatusEnumRejected TransactionStatusEnum = "rejected"
)

func (e *TransactionStatusEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionStatusEnum(s)
	case string:
		*e = TransactionStatusEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionStatusEnum: %T", src)
	}
	return nil
}

type CardTier struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Tier int32  `json:"tier"`
}

type CreditRequest struct {
	ReferenceNumber     int64                 `json:"reference_number"`
	UserID              int32                 `json:"user_id"`
	Program             int32                 `json:"program"`
	MemberID            string                `json:"member_id"`
	TransactionTime     sql.NullTime          `json:"transaction_time"`
	CreditUsed          float64               `json:"credit_used"`
	RewardShouldReceive float64               `json:"reward_should_receive"`
	PromoUsed           sql.NullInt32         `json:"promo_used"`
	TransactionStatus   TransactionStatusEnum `json:"transaction_status"`
}

type LoyaltyMembership struct {
	ID      int64  `json:"id"`
	Program int32  `json:"program"`
	Name    string `json:"name"`
}

type LoyaltyProgram struct {
	ID                 int64          `json:"id"`
	Name               string         `json:"name"`
	CurrencyName       string         `json:"currency_name"`
	ProcessingTime     string         `json:"processing_time"`
	Description        sql.NullString `json:"description"`
	EnrollmentLink     string         `json:"enrollment_link"`
	TermsConditionLink string         `json:"terms_condition_link"`
	FormatRegex        string         `json:"format_regex"`
	PartnerCode        string         `json:"partner_code"`
	InitialEarnRate    float64        `json:"initial_earn_rate"`
}

type Promotion struct {
	ID                int64            `json:"id"`
	Program           int32            `json:"program"`
	PromoType         PromoTypeEnum    `json:"promo_type"`
	StartDate         time.Time        `json:"start_date"`
	EndDate           time.Time        `json:"end_date"`
	EarnRateType      EarnRateTypeEnum `json:"earn_rate_type"`
	Constant          float64          `json:"constant"`
	CardTier          sql.NullInt32    `json:"card_tier"`
	LoyaltyMembership sql.NullInt32    `json:"loyalty_membership"`
}

type User struct {
	ID            int64          `json:"id"`
	FullName      sql.NullString `json:"full_name"`
	CreditBalance float64        `json:"credit_balance"`
	Email         string         `json:"email"`
	Contact       sql.NullInt32  `json:"contact"`
	Password      string         `json:"password"`
	UserName      string         `json:"user_name"`
	CardTier      sql.NullInt32  `json:"card_tier"`
	CreatedAt     sql.NullTime   `json:"created_at"`
}
