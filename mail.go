package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "autogradingsystem99999@gmail.com"
	password := "nedlsjaxafqmlnms"
	to := []string{
		"jaiminparmar99999@gmail.com",
	}
	// Receiver email address.
	Receivermail := "jaiminparmar99999@gmail.com"
	subjet := "Email FORGOT"
	body := "this is body inclue mail.go"
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("To: " + Receivermail + "\r\n" +
		"Subject: " + subjet + "!\r\n" +
		"\r\n" +
		body + "\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
