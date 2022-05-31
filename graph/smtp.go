package graph

import (
	"errors"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(to []string) error {

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"

	// This is the message to send in the mail
	msg := "Hello World!"

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, to, body)

	if err != nil {
		return errors.New("Sending email failed")
	}

	return err
}
