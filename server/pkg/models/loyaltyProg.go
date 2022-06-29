package models

import (
  "gorm.io/gorm"
  "esc/ascendaRoyaltyPoint/pkg/config"

)

var db *gorm.DB

type LoyaltyProgram struct{
  gorm.Model
  ID int64 `gorm:"primaryKey`;
  ProgramName string `json:programName`
  CurrencyName string `json:currencyName`
  ProcessTime string `json:processingTime`
  Description string `json:description`
  EnrollmentLink string `json:enrollmentLink`
  TermsConditionLink string `json:termsConditionLink`
  FormatRegEx string `json:formatRegEx`
  Partnercode string `json:partnerCode`
  InitialEarnRate float32 `json:initialEarnRate`
}

func init(){
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&LoyaltyProgram{})
}

func (p *LoyaltyProgram) CreateLoyaltyProgram() *LoyaltyProgram{
  db.Create(&p)
  return p
}

func GetAllProg() []LoyaltyProgram{
  var progs []LoyaltyProgram
  db.Find(&progs)
  return progs
}

func GetProgById(ProgId int64) (*LoyaltyProgram, *gorm.DB){
  var getProg LoyaltyProgram
  db := db.Where("ID=?",ProgId).Find(&getProg)
  return &getProg,db 
}

func DeleteProg(ProgId int64 ) LoyaltyProgram{
  var prog LoyaltyProgram
  db.Where("ID=?",ProgId).Delete(prog)
  return prog
}
