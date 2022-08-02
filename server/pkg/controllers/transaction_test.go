package controllers

import (
	"context"
	"database/sql"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"esc/ascendaRoyaltyPoint/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createLoyaltyObject() models.CreateLoyaltyParams {
	arg := models.CreateLoyaltyParams{
		Name:               utils.RandomString(6),
		CurrencyName:       utils.RandomString(6),
		ProcessingTime:     utils.RandomString(4),
		Description:        utils.RandomString(20),
		EnrollmentLink:     utils.RandomString(20),
		TermsConditionLink: utils.RandomString(20),
		FormatRegex:        "100\\d{4}$",
		PartnerCode:        utils.RandomString(5),
		InitialEarnRate:    utils.RandomFloat(10),
	}
	return arg
}
func createUserObject(carTier sql.NullInt32) models.CreateUserParams {
	arg := models.CreateUserParams{
		FullName:      sql.NullString{Valid: true, String: utils.RandomString(6)},
		CreditBalance: 2000,
		Email:         utils.RandomString(6),
		Contact:       sql.NullInt32{Valid: false},
		Password:      utils.RandomString(10),
		UserName:      utils.RandomString(5),
		CardTier:      carTier,
	}
	return arg
}
func createCardTierObject(carTier int32) models.CreateCardTierParams {
	arg := models.CreateCardTierParams{
		Name: utils.RandomString(6),
		Tier: carTier,
	}
	return arg
}

// add with onetime promo: kangming
// multiple promotion :kangming
// multiply onetime promo :nicholas
// multiply ongoing:nicholas
// add when there is a cardtier requirement for promo:kangming
// multiply when there is a cardtier requirement for promo:nicholas
// promo date range outside today's date

// when promo ask for cartier but use no cardtier/ user's cardtier is below what is requested
func TestRewardCalPromoOutOfRange(t *testing.T) {

	cardTierArgs := models.CreateCardTierParams{
		Name: utils.RandomString(7),
		Tier: 2,
	}
	cardTier, err := testQueries.CreateCardTier(context.Background(), cardTierArgs)
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-07-15")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
		CardTier:     sql.NullInt32{Valid: true, Int32: int32(cardTier.ID)},
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, expected, result)
}

// when promo ask for cartier but use no cardtier/ user's cardtier is below what is requested
func TestRewardCalCardTier(t *testing.T) {

	cardTierArgs := models.CreateCardTierParams{
		Name: utils.RandomString(7),
		Tier: 2,
	}
	cardTier, err := testQueries.CreateCardTier(context.Background(), cardTierArgs)
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
		CardTier:     sql.NullInt32{Valid: true, Int32: int32(cardTier.ID)},
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, expected, result)

}

//  when no promo
func TestRewardCalAddNoPromo(t *testing.T) {

	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, expected, result)

}

func TestRewardCalNormal(t *testing.T) {
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer*(program.InitialEarnRate) + constant
	require.Equal(t, expected, result)

}

// add with one time promo, when promo is not used
func TestAddOneTimePromo(t *testing.T) {
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}

	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("onetime"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
	}

	createCreditRequestArgs := models.CreateCreditRequestParams{
		UserID:              int32(user.ID),
		PromoUsed:           sql.NullInt32{Valid: false},
		Program:             int32(program.ID),
		MemberID:            utils.RandomString(20),
		TransactionTime:     sql.NullTime{Valid: true, Time: startDate},
		CreditUsed:          creditToTransfer,
		RewardShouldReceive: utils.RandomFloat(30),
		TransactionStatus:   models.TransactionStatusEnum("pending"),
	}

	_, err = testQueries.CreateCreditRequest(context.Background(), createCreditRequestArgs)
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, expected, result)
}

// Test for Mul One Time Promotion with Previous Credit Request
func TestMulOneTimePromoWithCreditRequest(t *testing.T) {
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("onetime"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("mul"),
		Constant:     float64(constant),
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, result, expected)
}

// test when there is multiple promotion going on, take the one with greatest return
func TestMultiplePromotions(t *testing.T) {
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: creditToTransfer,
		MembershipId:     utils.RandomString(6),
	}

	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant1 float64 = 1000
	createPromoArgs1 := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant1),
	}

	var constant2 float64 = 200
	createPromoArgs2 := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("mul"),
		Constant:     float64(constant2),
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs1)
	require.NoError(t, err)
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs2)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate) * constant2
	require.Equal(t, result, expected)
}

// Test for Mul One Time Promotion without Previous Credit Request
// func TestMulOneTimePromoWithoutCreditRequest(t *testing.T) {
// 	createLoyaltyArgs := createLoyaltyObject()
// 	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
// 	var creditToTransfer float64 = 100
// 	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
// 	require.NoError(t, err)
// 	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
// 	require.NoError(t, err)

// 	args := models.TransferParams{
// 		UserId:           int32(user.ID),
// 		ProgramId:        int32(program.ID),
// 		CreditToTransfer: float64(creditToTransfer),
// 		MembershipId:     utils.RandomString(6),
// 	}
// 	startDate, err := time.Parse("2006-01-02", "2022-07-01")
// 	require.NoError(t, err)
// 	endDate, err := time.Parse("2006-01-02", "2022-07-30")
// 	require.NoError(t, err)

// 	var constant float64 = 1000
// 	createPromoArgs := models.CreatePromotionParams{
// 		Program:      int32(program.ID),
// 		PromoType:    models.PromoTypeEnum("onetime"),
// 		StartDate:    startDate,
// 		EndDate:      endDate,
// 		EarnRateType: models.EarnRateTypeEnum("mul"),
// 		Constant:     float64(constant),
// 	}

// 	createCreditRequestArgs := models.CreateCreditRequestParams{
// 		UserID:              int32(user.ID),
// 		PromoUsed:           sql.NullInt32{Valid: false},
// 		Program:             int32(program.ID),
// 		MemberID:            utils.RandomString(6),
// 		TransactionTime:     sql.NullTime{},
// 		CreditUsed:          creditToTransfer,
// 		RewardShouldReceive: utils.RandomFloat(10),
// 		TransactionStatus:   models.TransactionStatusEnumApproved,
// 	}
// 	_, err = testQueries.CreateCreditRequest(context.Background(), createCreditRequestArgs)
// 	require.NoError(t, err)
// 	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
// 	require.NoError(t, err)
// 	result, _, err := CalculateReward(context.Background(), testQueries, args)
// 	require.NoError(t, err)
// 	expected := creditToTransfer * (program.InitialEarnRate) * constant
// 	require.Equal(t, result, expected)
// }

// Test for Mul Ongoing Promotion that Requires Card Tier and User's cardtier matches
func TestMulRequireCardTierMatch(t *testing.T) {
	createCardTierArgs := createCardTierObject(1)
	card_tier, err := testQueries.CreateCardTier(context.Background(), createCardTierArgs)
	require.NoError(t, err)

	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Int32: int32(card_tier.ID), Valid: true})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)

	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("mul"),
		Constant:     float64(constant),
		CardTier:     (sql.NullInt32{Int32: int32(card_tier.ID), Valid: true}),
	}

	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate) * constant
	require.Equal(t, result, expected)
}

// add when there is a card tier requirement for promo, when user card tier == promo card tier
func TestAddCardTier(t *testing.T) {
	createCardTierArgs := models.CreateCardTierParams{
		Name: utils.RandomString(6),
		Tier: 3,
	}
	cardTier, err := testQueries.CreateCardTier(context.Background(), createCardTierArgs)
	require.NoError(t, err)

	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: true, Int32: int32(cardTier.ID)})

	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: creditToTransfer,
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)
	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
		CardTier:     sql.NullInt32{Valid: true, Int32: int32(cardTier.ID)},
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer*(program.InitialEarnRate) + constant
	require.Equal(t, result, expected)
}

// Test for Mul Ongoing Promotion
func TestMulOnGoingPromo(t *testing.T) {
	createLoyaltyArgs := createLoyaltyObject()
	createUserArgs := createUserObject(sql.NullInt32{Valid: false})
	var creditToTransfer float64 = 100
	program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), createUserArgs)
	require.NoError(t, err)
	args := models.TransferParams{
		UserId:           int32(user.ID),
		ProgramId:        int32(program.ID),
		CreditToTransfer: float64(creditToTransfer),
		MembershipId:     utils.RandomString(6),
	}
	startDate, err := time.Parse("2006-01-02", "2022-07-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2022-08-30")
	require.NoError(t, err)

	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("mul"),
		Constant:     float64(constant),
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, _, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate) * constant
	require.Equal(t, result, expected)
}
