package main

import (
	"crypto/tls"
	"os"

	gomail "gopkg.in/mail.v2"
)

//--------------- (EMAIL HANDLER) ----------------
func email_handler(user_name string) {

	InitEnvFile()

	aws_secrets := AwsInit()

	email := gomail.NewMessage()

	email_send := aws_secrets["EMAIL_SENDER"]
	email_password := aws_secrets["EMAIL_SENDER_PASSWORD"]
	email_subject := "AWS securty - Rotate access key"
	email_body := "Greetings user,\n\nYou are receiving this email as a gentle reminder to decommission your existing 'ACCESS-KEY' in the 'NowFloats-DEV' account as it is older than 90 days. This is currently above the proposed compliance limit.\n\nWe trust you will do the needful at the earliest."
	email_connection := os.Getenv("EMAIL_CONNECTION")
	port := 587

	email.SetHeader("From", email_send)
	email.SetHeader("To", user_name)
	email.SetHeader("Subject", email_subject)
	email.SetBody("text/plain", email_body)

	d := gomail.NewDialer(email_connection, port, email_send, email_password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(email)
	CheckForNil(err)
}
//------------ (END EMAIL HANDLER) --------------