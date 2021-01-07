package helper

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"regexp"
)

//LogError log the error to info.log
func LogError(message string) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print(message)
}

//CheckEmail Check if the email is valid
func CheckEmail(email string) error {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(email) {
		return nil
	}
	return errors.New("Email is not correctly formated")
}

//SendEmail send mail
func SendEmail(reciver []string, message []byte) {
	// Sender data.
	from := "notifyuseruzcom@gmail.com"
	password := "thisIsFunToCreate0924"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, reciver, message)
	if err != nil {
		LogError(err.Error())
		log.Fatal(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
