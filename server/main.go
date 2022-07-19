package main

import (
  "log"
  "esc/ascendaRoyaltyPoint/pkg/controllers"
  "esc/ascendaRoyaltyPoint/pkg/config"
  "esc/ascendaRoyaltyPoint/pkg/models"
)

func main(){
  config.Connect()
  db := config.GetDB()
  store :=models.NewStore(db)
  server := controllers.NewServer(store)
  err := server.Start("0.0.0.0:8080")
  if err!=nil{
    log.Fatal("cannot start server",err)
  }
}

