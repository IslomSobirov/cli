package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strings"
)

//SMSRequestBody struct for sms service
type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

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

//SendSms send sms to clients
func SendSms(message string, phoneNumber string) {
	phoneNumber = strings.Replace(phoneNumber, "+", "", 1)
	body := SMSRequestBody{
		APIKey:    os.Getenv("NEXMO_API_KEY"),
		APISecret: os.Getenv("NEXMP_API_SECRET"),
		To:        phoneNumber,
		From:      "Store",
		Text:      message,
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		LogError(err.Error())
		log.Fatal(err.Error())
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		LogError(err.Error())
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogError(err.Error())
		log.Fatal(err.Error())
	}

	fmt.Println(string(respBody))
}
