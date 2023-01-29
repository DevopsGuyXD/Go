package main

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/mail.v2"
)

func main() {

	email := gomail.NewMessage()

  email_subject := "AWS securty - Rotate access key"
  email_body := "Greetings user,\nYou are receiving this email as a gentle reminder to decommission your existing ACCESS KEY in the AWS GETKITSUNE account as it is now older than 90 days. This is currently above the proposed compliance limit.\nWe trust you will do the needful at the earliest."
  email_from := "nfcloudsecurity@nowfloats.com"
  email_password := "P67%4urG123"
  email_to := "bharath.dundi@nowfloats.com"
  smtp_protocal := "smtp.gmail.com"
  port := 587

  email.SetHeader("From", email_from)
  email.SetHeader("To", email_to)
  email.SetHeader("Subject", email_subject)
  email.SetBody("text/plain", email_body)

  d := gomail.NewDialer(smtp_protocal, port, email_from, email_password)
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  err := d.DialAndSend(email); if err != nil{
    log.Fatal(err)
  }
}