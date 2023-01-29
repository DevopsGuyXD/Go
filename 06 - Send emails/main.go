package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
  m := gomail.NewMessage()

  m.SetHeader("From", "nfcloudsecurity@nowfloats.com")

  m.SetHeader("To", "bharath.dundi@nowfloats.com")

  m.SetHeader("Subject", "Gomail test subject")

  m.SetBody("text/plain", "This is Gomail test body")

  d := gomail.NewDialer("smtp.gmail.com", 587, "nfcloudsecurity@nowfloats.com", "P67%4urG123")

  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  // Now send E-Mail
  if err := d.DialAndSend(m); err != nil {
    fmt.Println(err)
    panic(err)
  }

  return
}