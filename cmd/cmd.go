package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/cli/helper"
	"github.com/manifoldco/promptui"
)

type client struct {
	id          int
	phoneNumber string
	email       string
}

type purchase struct {
	id    int
	name  string
	price int
}

var validateID = func(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		helper.LogError("ID is invalid: " + input)
		return errors.New("Invalid number")
	}
	return nil
}

//Init the project
func Init() {
	fmt.Println("Good morning this is sms/email sending service")
	//Select option to notify client
	prompt := promptui.Select{
		Label: "Select one option",
		Items: []string{"Sms", "Email"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	switch result {
	case "Sms":
		c := getClient(true)
		p := getPurchase()
		helper.SendSms(
			"Purchase: "+p.name+"\nPrice: "+strconv.Itoa(p.price)+" so'm",
			c.phoneNumber,
		)
	case "Email":

		c := getClient(false)
		p := getPurchase()

		message := []byte("To: " + c.email + "\r\n" +
			"Subject: Purchase has been accepted!\r\n" +
			"\r\n" +
			"Id: " + strconv.Itoa(p.id) + " Total price " + strconv.Itoa(p.price) + " so'm")

		helper.SendEmail([]string{
			c.email,
		}, message)

	}
}

func getClient(phoneNumber bool) client {

	//Get the id of client
	ClientIDPrompt := promptui.Prompt{
		Label:    "Please provide id of the client",
		Validate: validateID,
	}
	clientID, err := ClientIDPrompt.Run()
	clientIDD, _ := strconv.Atoi(clientID)

	if err != nil {
		helper.LogError(err.Error())
		log.Fatal(err.Error())
	}
	if phoneNumber {
		ClientNumber := promptui.Prompt{
			Label: "Please provide the phoneNumber of the client format: +998998403675",
		}

		clientNumber, err := ClientNumber.Run()
		if err != nil {
			helper.LogError(err.Error())
			log.Fatal(err.Error())
		}
		return client{
			id:          clientIDD,
			phoneNumber: clientNumber,
		}
	}
	//Get the email of client
	ClientEmail := promptui.Prompt{
		Label: "Please provide the email of the client",
		Validate: func(input string) error {
			err := helper.CheckEmail(input)
			if err != nil {
				helper.LogError("Email is invalid : " + input)
				return errors.New(err.Error())
			}
			return nil
		},
	}

	clientEmail, emailerr := ClientEmail.Run()
	if emailerr != nil {
		log.Fatal(emailerr.Error())
	}

	return client{
		id:    clientIDD,
		email: clientEmail,
	}

}

func getPurchase() purchase {
	//Get the id of purchase
	purchaseIDPrompt := promptui.Prompt{
		Label:    "Please provide id of the purchase",
		Validate: validateID,
	}

	purchaseID, purchaseIDError := purchaseIDPrompt.Run()
	purchaseIDD, _ := strconv.Atoi(purchaseID)
	if purchaseIDError != nil {
		helper.LogError(purchaseIDError.Error())
		log.Fatal(purchaseIDError)
	}

	//Get the purchase name
	purchaseNamePrompt := promptui.Prompt{
		Label: "Please enter the name of purchase",
		Validate: func(input string) error {
			if len(input) < 3 {
				helper.LogError("The name of the purchase has to be more than 3 characters. Ex:" + input)
				return errors.New("The name of the purchase has to be more than 3 characters")
			}

			return nil
		},
	}
	purchaseName, purchaseError := purchaseNamePrompt.Run()
	if purchaseError != nil {
		helper.LogError(purchaseError.Error())
		log.Fatal(purchaseError)
	}

	//Get the purchase price
	purchasePricePrompt := promptui.Prompt{
		Label: "Please enter the price of purchase",
		Validate: func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				helper.LogError(err.Error())
				return errors.New("price has to be in numbers only")
			}

			return nil
		},
	}
	purchasePrice, purchasePriceError := purchasePricePrompt.Run()
	price, _ := strconv.Atoi(purchasePrice)
	if purchasePriceError != nil {
		helper.LogError(purchasePriceError.Error())
		log.Fatal(purchasePriceError)
	}

	return purchase{
		id:    purchaseIDD,
		name:  purchaseName,
		price: price,
	}
}
