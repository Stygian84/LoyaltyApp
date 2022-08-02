package controllers

import (
	"context"
	"database/sql"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"esc/ascendaRoyaltyPoint/pkg/utils"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func FuzzCalRewards(f *testing.F) {

	f.Add(utils.RandomFloat(10000), utils.RandomFloat(10000))
	f.Fuzz(func(t *testing.T, n float64, n2 float64) {

		var constant float64 = n
		var base float64 = n2

		createLoyaltyArgs := createLoyaltyObject()
		program, err := testQueries.CreateLoyalty(context.Background(), createLoyaltyArgs)
		require.NoError(t, err)
		startDate, err := time.Parse("2006-01-02", "2022-07-01")
		require.NoError(t, err)
		endDate, err := time.Parse("2006-01-02", "2022-08-30")
		require.NoError(t, err)
		createPromoArgs := models.CreatePromotionParams{
			Program:      int32(program.ID),
			PromoType:    models.PromoTypeEnum("ongoing"),
			StartDate:    startDate,
			EndDate:      endDate,
			EarnRateType: models.EarnRateTypeEnum("mul"),
			Constant:     float64(constant),
		}

		promotion, err := testQueries.CreatePromotion(context.Background(), createPromoArgs)
		require.NoError(t, err)
		result := processReward(promotion, base)
		expected := base * promotion.Constant
		require.Equal(t, expected, result)

	})
}

func FuzzMulOnGoingPromo(f *testing.F) {
	f.Add(utils.RandomFloat(10000), utils.RandomFloat(10000))
	f.Fuzz(func(t *testing.T, n3 float64, n4 float64) {

		var constant float64 = math.Abs(n3)
		var creditToTransfer float64 = math.Abs(n4)

		createLoyaltyArgs := createLoyaltyObject()
		createUserArgs := createUserObject(sql.NullInt32{Valid: false})
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
		require.Equal(t, expected, result)
	})
}
