package controllers

import (
  "esc/ascendaRoyaltyPoint/pkg/models"
  "testing"
  "github.com/stretchr/testify/require"
  "context"
  "database/sql"
  "time"
  "esc/ascendaRoyaltyPoint/pkg/utils"
)
func createLoyaltyObject()models.CreateLoyaltyParams{
    arg := models.CreateLoyaltyParams{
      Name : utils.RandomString(6),
      CurrencyName:utils.RandomString(6),
      ProcessingTime:utils.RandomString(4),
      Description:sql.NullString{String:utils.RandomString(20),Valid:true},
      EnrollmentLink:utils.RandomString(20),
      TermsConditionLink:utils.RandomString(20),
      FormatRegex:"100\\d{4}$",
      PartnerCode:utils.RandomString(5),
      InitialEarnRate:utils.RandomFloat(10),
  }
  return arg
}
func createUserObject(carTier sql.NullInt32)models.CreateUserParams{
  arg:=models.CreateUserParams{
    FullName:sql.NullString{Valid:true,String:utils.RandomString(6)},
    CreditBalance:2000,
    Email:utils.RandomString(6),
    Contact:sql.NullInt32{Valid:false},
    Password:utils.RandomString(10),
    UserName:utils.RandomString(5),
    CardTier:carTier,
  }
  return arg
}
// add with onetime promo
// multiple promotion
// multiply onetime promo
// multiply ongoing
// add when no promo
// multiply when no promo
// add when there is a cardtier requirement for promo
// multiply when there is a cardtier requirement for promo
// when promo ask for cartier but use no cardtier/ user's cardtier is below what is requested




func TestRewardCalNormal(t *testing.T){
  createLoyaltyArgs:=createLoyaltyObject()
  createUserArgs:= createUserObject(sql.NullInt32{Valid:false})
  var creditToTransfer float64=100
  program,err:=testQueries.CreateLoyalty(context.Background(),createLoyaltyArgs)
  require.NoError(t,err)
  user,err := testQueries.CreateUser(context.Background(),createUserArgs)
  require.NoError(t,err)
  args:= models.TransferParams{
    UserId: int32(user.ID),
    ProgramId:int32(program.ID),
    CreditToTransfer:float64(creditToTransfer),
    MembershipId:utils.RandomString(6),
  }
  startDate,err:=time.Parse("2006-01-02","2022-07-01")
  require.NoError(t,err)
  endDate,err:=time.Parse("2006-01-02","2022-07-30")
  require.NoError(t,err)

  var constant float64=1000
  createPromoArgs:=models.CreatePromotionParams{
    Program:int32(program.ID),
    PromoType:models.PromoTypeEnum("ongoing"),
    StartDate:startDate,
    EndDate:endDate,
    EarnRateType:models.EarnRateTypeEnum("add"),
    Constant:float64(constant),
  }
  _ ,err=testQueries.CreatePromotion(context.Background(),createPromoArgs)
  require.NoError(t,err)
  result ,err:=CalculateReward(context.Background(),testQueries,args)
  require.NoError(t,err)
  expected := creditToTransfer*(program.InitialEarnRate)+constant 
  require.Equal(t,result,expected)

}

