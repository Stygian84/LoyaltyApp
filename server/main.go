package main

import (
	"esc/ascendaRoyaltyPoint/pkg/config"
	// "esc/ascendaRoyaltyPoint/pkg/handle"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"esc/ascendaRoyaltyPoint/pkg/controllers"
	"log"
)

func main() {
	config.Connect()
	db := config.GetDB()
	store := models.NewStore(db)
	server := controllers.NewServer(store)
	err := server.Start("0.0.0.0:8080")
	if err!=nil{
	  log.Fatal("cannot start server",err)
	}
	// err := handle.ReadHandbackFile()
	// err := handle.SendAccrual()
	// log.Fatal(err)
}
