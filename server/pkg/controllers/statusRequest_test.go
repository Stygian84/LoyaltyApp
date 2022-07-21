package controllers

import (
	"context"
	"database/sql"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"esc/ascendaRoyaltyPoint/pkg/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

// test credit request db function, one credit request under user
func TestQueryTransactionStatus(t *testing.T) {
	var expected []models.CreditRequest
	// create user and loyalty prog
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	createLoyaltyArgs := createLoyaltyObject()
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)

	// create credit request
	createCreditRequestArgs := models.CreateCreditRequestParams{
		UserID:              int32(user.ID),
		PromoUsed:           sql.NullInt32{Valid: false},
		Program:             int32(program.ID),
		MemberID:            utils.RandomString(6),
		TransactionTime:     sql.NullTime{Valid: false},
		CreditUsed:          creditToTransfer,
		RewardShouldReceive: creditToTransfer + 5,
		TransactionStatus:   models.TransactionStatusEnum("created"),
	}
	creditRequest, err := testQueries.CreateCreditRequest(context.Background(), createCreditRequestArgs)
	require.NoError(t, err)
	expected = append(expected, creditRequest)
	result, err := testQueries.GetCreditRequestByUser(context.Background(), int32(user.ID))
	require.NoError(t, err)
	require.Equal(t, result, expected)

	// test get transaction status function
	var expected2 []TransStatus
	expected2 = append(expected2, TransStatus{CreditRequestId: int32(creditRequest.ReferenceNumber), TransactionStatus: string(creditRequest.TransactionStatus), ProgramId: int32(program.ID)})
	result2 := GetTransStatus(result)
	require.Equal(t, result2, expected2)
}

// test credit request db function, no credit request under user
func TestQueryTransactionStatusNoRequest(t *testing.T) {
	var expected []models.CreditRequest
	// create user and loyalty prog
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	createLoyaltyArgs := createLoyaltyObject()
	_, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)

	expected = nil
	result, err := testQueries.GetCreditRequestByUser(context.Background(), int32(user.ID))
	require.NoError(t, err)
	require.Equal(t, result, expected)

}

// test credit requesr db function, multiple credit request under user
func TestQueryTransactionStatusMultipleRequest(t *testing.T) {
	var expected []models.CreditRequest
	// create user and loyalty prog
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	createLoyaltyArgs := createLoyaltyObject()
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)

	// create credit request
	createCreditRequestArgs := models.CreateCreditRequestParams{
		UserID:              int32(user.ID),
		PromoUsed:           sql.NullInt32{Valid: false},
		Program:             int32(program.ID),
		MemberID:            utils.RandomString(6),
		TransactionTime:     sql.NullTime{Valid: false},
		CreditUsed:          creditToTransfer,
		RewardShouldReceive: creditToTransfer + 5,
		TransactionStatus:   models.TransactionStatusEnum("created"),
	}
	creditRequest, err := testQueries.CreateCreditRequest(context.Background(), createCreditRequestArgs)
	require.NoError(t, err)

	// create credit request
	createCreditRequestArgs2 := models.CreateCreditRequestParams{
		UserID:              int32(user.ID),
		PromoUsed:           sql.NullInt32{Valid: false},
		Program:             int32(program.ID),
		MemberID:            utils.RandomString(6),
		TransactionTime:     sql.NullTime{Valid: false},
		CreditUsed:          creditToTransfer,
		RewardShouldReceive: creditToTransfer + 5,
		TransactionStatus:   models.TransactionStatusEnum("pending"),
	}
	creditRequest2, err := testQueries.CreateCreditRequest(context.Background(), createCreditRequestArgs2)
	require.NoError(t, err)
	expected = append(expected, creditRequest)
	expected = append(expected, creditRequest2)
	result, err := testQueries.GetCreditRequestByUser(context.Background(), int32(user.ID))
	require.NoError(t, err)
	require.Equal(t, result, expected)

	// test get transaction status function
	var expected2 []TransStatus
	expected2 = append(expected2, TransStatus{CreditRequestId: int32(creditRequest.ReferenceNumber), TransactionStatus: string(creditRequest.TransactionStatus), ProgramId: int32(program.ID)})
	expected2 = append(expected2, TransStatus{CreditRequestId: int32(creditRequest2.ReferenceNumber), TransactionStatus: string(creditRequest2.TransactionStatus), ProgramId: int32(program.ID)})
	result2 := GetTransStatus(result)
	require.Equal(t, result2, expected2)
}
