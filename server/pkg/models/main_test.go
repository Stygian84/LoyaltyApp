package models

import (
  "os"
  "testing"
  _ "github.com/lib/pq"
  "esc/ascendaRoyaltyPoint/pkg/config"
)

var testQueries *Queries
func TestMain(m *testing.M){

  config.Connect()
  testQueries = New(config.GetDB()) 
  os.Exit(m.Run())
}
