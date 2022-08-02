package main

import (
	"esc/ascendaRoyaltyPoint/pkg/config"
	"esc/ascendaRoyaltyPoint/pkg/handle"
	// "esc/ascendaRoyaltyPoint/pkg/handle"
	"esc/ascendaRoyaltyPoint/pkg/controllers"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"log"
)

func main() {

	go handle.RunCron("01:00", "03:00")

	config.Connect()
	db := config.GetDB()
	store := models.NewStore(db)
	server := controllers.NewServer(store)
	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
