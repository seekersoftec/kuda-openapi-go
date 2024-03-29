package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/seekersoftec/kuda-openapi-go/kuda"
)

type FakeData struct {
	FirstName         string `faker:"first_name"`
	LastName          string `faker:"last_name"`
	MiddleName        string `faker:"word"`
	Email             string `faker:"email"`
	PhoneNumber       string `faker:"e_164_phone_number"`
	BusinessName      string `faker:"word"`
	TrackingReference int    `faker:"oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 27, 61, 13"`
	AccountNumber     string `faker:"-"`
}

func normalizePhoneNumber(phoneNumber string, numberLength int) (string, error) {
	// Replace the plus sign with zero
	normalizedNumber := strings.ReplaceAll(phoneNumber, "+", "0")

	// Remove the first three digits and replace them with a random value
	rand.NewSource(time.Now().UnixNano())
	replacement := []string{"080", "081", "070", "090"}
	if len(normalizedNumber) > 3 {
		normalizedNumber = replacement[rand.Intn(len(replacement))] + normalizedNumber[3:]
	}

	// Check if the resulting string is N digits long
	if len(normalizedNumber) < numberLength {
		normalizedNumber = fmt.Sprintf("%-*s", numberLength, normalizedNumber)
	} else if len(normalizedNumber) > numberLength {
		normalizedNumber = normalizedNumber[:numberLength]
	}
	return normalizedNumber, nil
}

// Single fake function can be used for retrieving particular values.
func generateFakeData() (FakeData, error) {
	var generatedFakeData FakeData
	err := faker.FakeData(&generatedFakeData)
	if err != nil {
		return FakeData{}, err
	}

	normalizedNumber, err := normalizePhoneNumber(generatedFakeData.PhoneNumber, 11)
	if err != nil {
		log.Fatalf("Error normalizing phone number: %s", err)
	}
	// fmt.Printf("Normalized phone number: %s", normalizedNumber)
	generatedFakeData.PhoneNumber = normalizedNumber

	// fmt.Printf("%+v", generatedFakeData)
	return generatedFakeData, nil
}

func main() {
	connectionParams := &kuda.ConnectionParams{
		Email:       "seekersoftec@gmail.com", // "test@example.com",
		ApiKey:      "6bD9FeXrKMZq4zndfuU1",   //"api-key",
		Live:        false,
		ShowRequest: false,
		ShowHeader:  false,
		BaseURL:     "",
	}

	k, err := kuda.NewKudaBankAPI(connectionParams)
	if err != nil {
		log.Fatalf("Errorf: %s ", err)
	}

	// Get token
	// token, err := k.GetToken()
	// if err != nil {
	// 	log.Fatalf("Errof: %s ", err)
	// }
	// fmt.Println(token)

	// Name Enquiry
	beneficiaryAccountNumber := "2611457591"
	beneficiaryBankCode := "999057"

	name, err := k.NameEnquiry(beneficiaryAccountNumber, beneficiaryBankCode, "", "")
	if err != nil {
		log.Fatalf("Errof: %s ", err)
	}
	fmt.Println("NameEnquiry => ", name)

	// list of banks
	// banks, err := k.BankList()
	// if err != nil {
	// 	log.Fatalf("Errof: %s ", err)
	// }

	// fmt.Println(banks)

	/*
		Create Personas
	*/

	users := make(map[int]FakeData)
	count := 1
	for i := 0; i < count; i++ {
		fdata, err := generateFakeData()
		if err != nil {
			log.Fatalf("Errof: %s ", err)
		}
		fdata.TrackingReference = k.GetRef(10)
		fmt.Println(fdata)

		// Create a Virtual Accounts
		trackingRef := strconv.Itoa(fdata.TrackingReference)
		data, err := k.CreateVirtualAccount(fdata.FirstName, fdata.LastName, fdata.MiddleName, fdata.Email, fdata.PhoneNumber, fdata.BusinessName, trackingRef)
		if err != nil {
			log.Fatalf("Errof: %s ", err)
		}

		if strings.Contains(data["message"].(string), "Request successful.") && data["status"] == true {
			response := data["data"].(map[string]interface{})
			fmt.Println(response["accountNumber"])
			// store users
			fdata.AccountNumber = response["accountNumber"].(string)
			users[i] = fdata
		} else {
			fmt.Println(data)
		}
		// users[i] = fdata
	}
	//
	fmt.Println(users)
	//
	// SingleFundTransfer
	// Set test parameters
	// user1 := users[0]
	user2 := users[0]
	fmt.Println(user2)
	trackingRef := "5518871192" // strconv.Itoa(user2.TrackingReference) // user2.TrackingReference
	beneficiaryAccountNumber = user2.AccountNumber
	// beneficiaryAccountNumber, err := strconv.Atoi(user2.AccountNumber)
	// if err != nil {
	// handle error
	// 	log.Fatalf("Errof: %s ", err)
	// }
	beneficiaryBankCode = "999057"
	beneficiaryName := user2.FirstName + " " + user2.LastName
	amount := 500
	senderName := "Kuda User"
	narration := "Test transfer"
	ClientAccountNumber := "2012017027"
	// ClientAccountNumber, err := strconv.Atoi("2012017027") // 3000662179, 2012017027, user1.AccountNumber
	// if err != nil {
	// // handle error
	// log.Fatalf("Errof: %s ", err)
	// }
	nameEnquirySessionID := ""

	// SingleFundTransfer
	data, err := k.SingleFundTransfer(trackingRef, beneficiaryAccountNumber, beneficiaryBankCode, beneficiaryName, amount, senderName, narration, ClientAccountNumber, nameEnquirySessionID)
	if err != nil {
		log.Fatalf("Errof: %s ", err)
	}
	fmt.Println("SingleFundTransfer => ", data)

	// Virtual Fund Transfer
	data, err = k.VirtualFundTransfer(trackingRef, beneficiaryAccountNumber, beneficiaryBankCode, beneficiaryName, senderName, nameEnquirySessionID, amount, narration)
	if err != nil {
		log.Fatalf("Errof: %s ", err)
	}
	fmt.Println("VirtualFundTransfer => ", data)

	// Withdraw Fund to main account
	data, err = k.WithdrawVirtualAccount(trackingRef, amount, narration)
	if err != nil {
		log.Fatalf("Errof: %s ", err)
	}
	fmt.Println("WithdrawVirtualAccount => ", data)
}
