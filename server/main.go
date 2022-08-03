package main

import (
	"esc/ascendaRoyaltyPoint/pkg/config"
	"esc/ascendaRoyaltyPoint/pkg/controllers"
	_"esc/ascendaRoyaltyPoint/pkg/handle"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"log"
)

// "esc/ascendaRoyaltyPoint/pkg/handle"

func main() {

	// go handle.RunCron("09:54", "09:55")

	config.Connect()
	db := config.GetDB()
	store := models.NewStore(db)
	server := controllers.NewServer(store)
	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
