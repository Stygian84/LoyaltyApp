package main

import (
	"net/smtp"
)

const (
	email    = "stygian8442@gmail.com"
	email_pw = "mvmeztlcrqqclfxc"
	email_to = "ghostkirito84@gmail.com"
)

func main() {
	// config.Connect()
	// db := config.GetDB()
	// store := models.NewStore(db)
	// server := controllers.NewServer(store)
	// err := server.Start("0.0.0.0:8080")
	// if err != nil {
	// 	log.Fatal("cannot start server", err)
	// }
	// handle.RunCron("01:00", "03:00")
	//handle.SendAccrual()
	sendEmail()
}

func sendEmail() {
	from := email
	password := email_pw

	toEmailAddress := email_to
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
