
package controllers

import (
  "os"
  "testing"
  _ "github.com/lib/pq"
  "esc/ascendaRoyaltyPoint/pkg/config"
  "esc/ascendaRoyaltyPoint/pkg/models"

)

var testQueries *models.Queries
func TestMain(m *testing.M){

  config.Connect()
  testQueries = models.New(config.GetDB()) 
  os.Exit(m.Run())
}
