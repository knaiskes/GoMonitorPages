package main

import (
	"log"
	"net/smtp"
)

func email(msg string) {

	// enter your credentials
	receiver := ""
	sender := ""
	passw := ""

	subject := "Subject: Site's status code"
	body := msg

	auth := smtp.PlainAuth(
		"",
		sender,
		passw,
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		receiver,
		[]string{receiver},
		[]byte(subject+"\r\n"+body+"\r\n"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
