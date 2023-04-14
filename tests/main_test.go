package tests

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/seekersoftec/kuda-openapi-go/kuda"
)

var params = &kuda.ConnectionParams{
	Email:       "seekersoftec@gmail.com", // "test@example.com",
	ApiKey:      "ZAVRrek5uQEf0XtvmaBY",   //"api-key",
	Live:        false,
	ShowRequest: false,
	ShowHeader:  false,
	BaseURL:     "",
}

type FakeData struct {
	FirstName         string `faker:"first_name"`
	LastName          string `faker:"last_name"`
	MiddleName        string `faker:"word"`
	Email             string `faker:"email"`
	PhoneNumber       string `faker:"e_164_phone_number"`
	BusinessName      string `faker:"word"`
	TrackingReference string `faker:"oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 27, 61, 13"`
	AccountNumber     string `faker:"-"`
}

type Users map[int]FakeData

func SetUpKudaBankAPI(connectionParams *kuda.ConnectionParams) *kuda.KudaBankAPI {
	/*
		Setup Kuda Bank API
	*/
	api, err := kuda.NewKudaBankAPI(connectionParams)
	if err != nil {
		log.Fatalf("Error calling BankList: %s", err.Error())
	}
	return api
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

func CreateUsers(numOfUsers int) Users {
	/*
		Create Personas
	*/
	users := make(Users)
	for i := 0; i < numOfUsers; i++ {
		fdata, err := generateFakeData()
		if err != nil {
			log.Fatalf("Errof: %s ", err)
		}
		// fmt.Println(fdata)
		//
		users[i] = fdata
	}
	//
	// fmt.Println(users)

	return users
}
