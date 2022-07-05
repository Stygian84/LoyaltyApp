package models

import (
  "testing"
  "github.com/stretchr/testify/require"
  "database/sql"
  "context"
  "esc/ascendaRoyaltyPoint/pkg/utils"
)
func createLoyaltyObject()CreateLoyaltyParams{
    arg := CreateLoyaltyParams{
      Name : utils.RandomString(6),
      CurrencyName:utils.RandomString(6),
      ProcessingTime:utils.RandomString(4),
      Description:sql.NullString{String:utils.RandomString(20),Valid:true},
      EnrollmentLink:utils.RandomString(20),
      TermsConditionLink:utils.RandomString(20),
      FormatRegex:utils.RandomString(10),
      PartnerCode:utils.RandomString(5),
      InitialEarnRate:utils.RandomFloat(10),
  }
  return arg
}
func TestCreateLoyalty(t *testing.T){
  obj := createLoyaltyObject()
  program,err := testQueries.CreateLoyalty(context.Background(),obj)
  require.NoError(t,err)
  require.NotEmpty(t,program)

  require.Equal(t, program.Name,obj.Name)
  require.Equal(t, program.CurrencyName,obj.CurrencyName)
  require.Equal(t, program.ProcessingTime,obj.ProcessingTime)
  require.Equal(t, program.Description.String,obj.Description.String)
  require.Equal(t, program.EnrollmentLink,obj.EnrollmentLink)
  require.Equal(t, program.TermsConditionLink,obj.TermsConditionLink)
  require.Equal(t, program.FormatRegex,obj.FormatRegex)
  require.Equal(t, program.PartnerCode,obj.PartnerCode)
  require.Equal(t, program.InitialEarnRate,obj.InitialEarnRate)
  require.NotZero(t,program.ID)
}

func TestGetLoyaltyByID(t *testing.T){
  obj := createLoyaltyObject()
  program,err := testQueries.CreateLoyalty(context.Background(),obj)
  require.NoError(t,err)
  require.NotEmpty(t,program)

  require.NotZero(t,program.ID)
  foundProg,err := testQueries.GetLoyaltyByID(context.Background(),program.ID)
  require.Equal(t, foundProg.Name,obj.Name)
  require.Equal(t, foundProg.CurrencyName,obj.CurrencyName)
  require.Equal(t, foundProg.ProcessingTime,obj.ProcessingTime)
  require.Equal(t, foundProg.Description.String,obj.Description.String)
  require.Equal(t, foundProg.EnrollmentLink,obj.EnrollmentLink)
  require.Equal(t, foundProg.TermsConditionLink,obj.TermsConditionLink)
  require.Equal(t, foundProg.FormatRegex,obj.FormatRegex)
  require.Equal(t, foundProg.PartnerCode,obj.PartnerCode)
  require.Equal(t, foundProg.InitialEarnRate,obj.InitialEarnRate)
  require.Equal(t, foundProg.ID,program.ID)
}
func TestGetLoyaltyByName(t *testing.T){
  obj := createLoyaltyObject()
  program,err := testQueries.CreateLoyalty(context.Background(),obj)
  require.NoError(t,err)
  require.NotEmpty(t,program)

  require.NotZero(t,program.ID)
  foundProg,err := testQueries.GetLoyaltyByName(context.Background(),program.Name)
  require.Equal(t, foundProg.Name,obj.Name)
  require.Equal(t, foundProg.CurrencyName,obj.CurrencyName)
  require.Equal(t, foundProg.ProcessingTime,obj.ProcessingTime)
  require.Equal(t, foundProg.Description.String,obj.Description.String)
  require.Equal(t, foundProg.EnrollmentLink,obj.EnrollmentLink)
  require.Equal(t, foundProg.TermsConditionLink,obj.TermsConditionLink)
  require.Equal(t, foundProg.FormatRegex,obj.FormatRegex)
  require.Equal(t, foundProg.PartnerCode,obj.PartnerCode)
  require.Equal(t, foundProg.InitialEarnRate,obj.InitialEarnRate)
  require.Equal(t, foundProg.ID,program.ID)
}

func TestListLoyalty(t *testing.T){
  for i:=0;i<10;i++{
    createLoyaltyObject()
  }
  programs,err := testQueries.ListLoyalty(context.Background())
  require.NoError(t,err)
  for _,program :=range programs{
    require.NotEmpty(t,program)
  }
}

func TestUpdateLoyaltyAllField(t *testing.T){
  obj := createLoyaltyObject()
  program,err := testQueries.CreateLoyalty(context.Background(),obj)
  require.NoError(t,err)
  require.NotEmpty(t,program)
  arg:=UpdateLoyaltyParams{
    Name:"update",
    CurrencyName:"updatedCureency",
    ProcessingTime:"Instant!",
    Description:sql.NullString{String:"hello world",Valid:true},
    EnrollmentLink:"SomeLink",
    TermsConditionLink:"Another Link",
    FormatRegex:"A regex",
    PartnerCode:"CODE",
    InitialEarnRate:2,
    ID:program.ID,
  }
  updatedProg,updateError := testQueries.UpdateLoyalty(context.Background(),arg)
  require.NoError(t,updateError)
  require.Equal(t,updatedProg.Name,arg.Name)
  require.Equal(t, updatedProg.CurrencyName,arg.CurrencyName)
  require.Equal(t, updatedProg.ProcessingTime,arg.ProcessingTime)
  require.Equal(t, updatedProg.Description.String,arg.Description.String)
  require.Equal(t, updatedProg.EnrollmentLink,arg.EnrollmentLink)
  require.Equal(t, updatedProg.TermsConditionLink,arg.TermsConditionLink)
  require.Equal(t, updatedProg.FormatRegex,arg.FormatRegex)
  require.Equal(t, updatedProg.PartnerCode,arg.PartnerCode)
  require.Equal(t, updatedProg.InitialEarnRate,arg.InitialEarnRate)
  require.Equal(t, updatedProg.ID,program.ID)
}

func TestDeleteLoyalty(t *testing.T){
  obj := createLoyaltyObject()
  program,err := testQueries.CreateLoyalty(context.Background(),obj)
  require.NoError(t,err)
  require.NotEmpty(t,program)
  
  deleteErr := testQueries.DeleteLoyalty(context.Background(),program.ID)
  require.NoError(t,deleteErr)

  _,getErr := testQueries.GetLoyaltyByID(context.Background(),program.ID)
  require.EqualError(t,getErr,"sql: no rows in result set")
}







