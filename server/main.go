package main

import "esc/ascendaRoyaltyPoint/pkg/handle"

func main() {

	// go handle.RunCron("15:24", "15:25")
	handle.SendAccrual()
	// config.Connect()
	// db := config.GetDB()
	// store := models.NewStore(db)
	// server := controllers.NewServer(store)
	// err := server.Start("0.0.0.0:8080")
	// if err != nil {
	// 	log.Fatal("cannot start server", err)
	// }

}
