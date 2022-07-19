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
		Description:        sql.NullString{String: utils.RandomString(20), Valid: true},
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

// add with onetime promo: kangming
// multiple promotion :kangming
// multiply onetime promo :nicholas
// multiply ongoing:nicholas
// add when no promo
// multiply when no promo
// add when there is a cardtier requirement for promo:kangming
// multiply when there is a cardtier requirement for promo:nicholas
// when promo ask for cartier but use no cardtier/ user's cardtier is below what is requested
// promo date range outside today's date

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
	endDate, err := time.Parse("2006-01-02", "2022-07-30")
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
	result, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer*(program.InitialEarnRate) + constant
	require.Equal(t, result, expected)

}

// add with one time promo
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
	endDate, err := time.Parse("2006-01-02", "2022-07-30")
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
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, result, expected)
}

// multiple promotion
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
	endDate, err := time.Parse("2006-01-02", "2022-07-30")
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
	result, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate) * constant2
	require.Equal(t, result, expected)
}

// add when there is a card tier requirement for promo
func TestAddCardTier(t *testing.T) {
	createCardTierArgs := models.CreateCardTierParams{
		Name: utils.RandomString(6),
		Tier: 3,
	}
	createCardTierArgs2 := models.CreateCardTierParams{
		Name: utils.RandomString(6),
		Tier: 3,
	}
	cardTier, err := testQueries.CreateCardTier(context.Background(), createCardTierArgs)
	require.NoError(t, err)
	cardTier2, err := testQueries.CreateCardTier(context.Background(), createCardTierArgs2)
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
	endDate, err := time.Parse("2006-01-02", "2022-07-30")
	require.NoError(t, err)
	var constant float64 = 1000
	createPromoArgs := models.CreatePromotionParams{
		Program:      int32(program.ID),
		PromoType:    models.PromoTypeEnum("ongoing"),
		StartDate:    startDate,
		EndDate:      endDate,
		EarnRateType: models.EarnRateTypeEnum("add"),
		Constant:     float64(constant),
		CardTier:     sql.NullInt32{Valid: true, Int32: int32(cardTier2.ID)},
	}
	_, err = testQueries.CreatePromotion(context.Background(), createPromoArgs)
	require.NoError(t, err)
	result, err := CalculateReward(context.Background(), testQueries, args)
	require.NoError(t, err)
	expected := creditToTransfer * (program.InitialEarnRate)
	require.Equal(t, result, expected)
}
