package tests

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	// test with empty baseURL
	k := SetUpKudaBankAPI(params)
	response, err := k.GetToken()
	if err != nil {
		t.Errorf("Errorf: Got %v", err)
	}

	// Check the response type
	ok := strings.Contains(response, "invalid")

	assert.Equal(t, true, !ok)
	// test with custom baseURL
	// k = SetUpKudaBankAPI() // "test@example.com", "api-key", true, false, false, "https://example.com")
	// assert.Equal(t, "https://example.com", k.url)
}

func TestGetRef(t *testing.T) {
	// setup kuda bank API
	k := SetUpKudaBankAPI(params)

	// test with length of 10
	ref := k.GetRef(10)
	assert.Equal(t, 10, len(strconv.Itoa(ref)))
}

func TestKoboToNGN(t *testing.T) {
	k := SetUpKudaBankAPI(params)

	// test with 10 kobo
	ngn := k.KoboToNGN(10)
	assert.Equal(t, 0.10, ngn)

	// test with 5000 kobo
	ngn = k.KoboToNGN(5000)
	assert.Equal(t, 50.00, ngn)
}

func TestNGNToKobo(t *testing.T) {
	k := SetUpKudaBankAPI(params)

	// test with 0.10 naira
	kobo := k.NGNToKobo(0.10)
	assert.Equal(t, 10, kobo)

	// test with 50 naira
	kobo = k.NGNToKobo(50)
	assert.Equal(t, 5000, kobo)
}

func TestBankList(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)

	// Call BankList method
	response, err := k.BankList()
	// Check for errors
	if err != nil {
		t.Errorf("Error calling BankList: %s", err.Error())
	}

	// Check the response type
	// if reflect.TypeOf(response) != reflect.TypeOf(kuda.JsonType{}) {
	// 	t.Errorf("Unexpected response type. Expected kuda.JsonType, got %v", reflect.TypeOf(response))
	// }
	// t.Log(response)
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestNameEnquiry(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)

	// Set test parameters
	accountNumber := "1234567890"
	bankCode := "011"
	senderTrackingReference := "123"

	// Call NameEnquiry method
	response, err := k.NameEnquiry(accountNumber, bankCode, senderTrackingReference, "")

	// Check for errors
	if err != nil {
		t.Errorf("Error calling NameEnquiry: %s", err.Error())
	}
	// t.Log(response)
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestCreateVirtualAccount(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(1)
	userA := users[0]

	// Call CreateVirtualAccount method
	userA.TrackingReference = strconv.Itoa(k.GetRef(10))
	response, err := k.CreateVirtualAccount(userA.FirstName, userA.LastName, userA.MiddleName, userA.Email, userA.PhoneNumber, userA.BusinessName, userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
	}

	// Check the response type
	// if reflect.TypeOf(response) != reflect.TypeOf(kuda.JsonType{}) {
	// 	t.Errorf("Unexpected response type. Expected kuda.JsonType, got %v", reflect.TypeOf(response))
	// }
	// t.Log(response)
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestGetSingleVirtualAccount(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(1)
	userA := users[0]

	// Call CreateVirtualAccount method
	userA.TrackingReference = strconv.Itoa(k.GetRef(10))
	response, err := k.CreateVirtualAccount(userA.FirstName, userA.LastName, userA.MiddleName, userA.Email, userA.PhoneNumber, userA.BusinessName, userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
	}
	status := response["status"].(bool)
	assert.Equal(t, true, status)

	// Call GetVirtualSingleAccount method
	response, err = k.GetSingleVirtualAccount(userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling GetVirtualSingleAccount: %s", err.Error())
	}

	// Check the response type
	// if reflect.TypeOf(response) != reflect.TypeOf(map[string]interface{}{}) {
	// 	t.Errorf("Unexpected response type. Expected map[string]interface{}, got %v", reflect.TypeOf(response))
	// }
	// t.Log(response)
	status = response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestGetVirtualAccountBalance(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(1)
	userA := users[0]

	// Call CreateVirtualAccount method
	userA.TrackingReference = strconv.Itoa(k.GetRef(10))
	response, err := k.CreateVirtualAccount(userA.FirstName, userA.LastName, userA.MiddleName, userA.Email, userA.PhoneNumber, userA.BusinessName, userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
	}
	status := response["status"].(bool)
	assert.Equal(t, true, status)

	// Call GetVirtualSingleAccount method
	response, err = k.GetVirtualAccountBalance(userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling GetVirtualSingleAccount: %s", err.Error())
	}

	// Check the response type
	// if reflect.TypeOf(response) != reflect.TypeOf(map[string]interface{}{}) {
	// 	t.Errorf("Unexpected response type. Expected map[string]interface{}, got %v", reflect.TypeOf(response))
	// }
	// t.Log(response)
	status = response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestSingleFundTransfer(t *testing.T) {
	// Initialize KudaBankAPI struct
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(2)

	for i := 0; i < len(users); i++ {
		//
		// Retrieve the struct value from the map
		user, ok := users[i]
		if !ok {
			// handle error
			t.Errorf("Error Retrieving the user for the map: %v", user)
		}

		// Call CreateVirtualAccount method
		data, err := k.CreateVirtualAccount(user.FirstName, user.LastName, user.MiddleName, user.Email, user.PhoneNumber, user.BusinessName, user.TrackingReference)

		// Check for errors
		if err != nil {
			t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
		}

		if data["message"] == "Request successful." && data["status"] == true {
			response := data["data"].(map[string]interface{})
			fmt.Println(response["accountNumber"])
			//
			// store users

			// Modify the field directly
			user.AccountNumber = response["accountNumber"].(string)

			// Store the modified struct value back in the map
			users[i] = user
		} else {

			// Modify the field directly
			user.AccountNumber = ""

			// Store the modified struct value back in the map
			users[i] = user
		}

	}

	//
	userA := users[0]
	userB := users[1]

	// Call SingleFundTransfer method
	trackingRef := userA.TrackingReference
	beneficiaryAccountNumber := userB.AccountNumber
	beneficiaryBankCode := "999057"
	beneficiaryName := userB.FirstName + " " + userB.LastName
	senderName := userA.FirstName + " " + userA.LastName
	nameEnquirySessionID := ""
	amount := 500
	narration := "Test Transfer"
	ClientAccountNumber := "2012017027"
	// k.SingleFundTransfer(userA.TrackingReference, userB.AccountNumber, "999057", userB.FirstName+" "+userB.LastName, 500, userA.FirstName+" "+userA.LastName, "Test Transfer", ClientAccountNumber, "")
	response, err := k.SingleFundTransfer(trackingRef, beneficiaryAccountNumber, beneficiaryBankCode, beneficiaryName, amount, senderName, narration, ClientAccountNumber, nameEnquirySessionID)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling SingleFundTransfer: %s", err.Error())
	}

	// t.Log(response)
	//
	if strings.Contains(response["message"].(string), "invalid") {
		// handle that
		t.Errorf("Error calling SingleFundTransfer: %s", response)
	}

	// t.Log(response)
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestVirtualFundTransfer(t *testing.T) {
	// Initialize API client with test server's URL
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(2)

	for i := 0; i < len(users); i++ {
		//
		// Retrieve the struct value from the map
		user, ok := users[i]
		if !ok {
			// handle error
			t.Errorf("Error Retrieving the user for the map: %v", user)
		}

		user.TrackingReference = strconv.Itoa(k.GetRef(10))
		// Call CreateVirtualAccount method
		data, err := k.CreateVirtualAccount(user.FirstName, user.LastName, user.MiddleName, user.Email, user.PhoneNumber, user.BusinessName, user.TrackingReference)

		// Check for errors
		if err != nil {
			t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
		}

		if data["message"] == "Request successful." && data["status"] == true {
			response := data["data"].(map[string]interface{})
			fmt.Println(response["accountNumber"])
			//
			// store users

			// Modify the field directly
			user.AccountNumber = response["accountNumber"].(string)

			// Fund account
			response, err := k.FundVirtualAccount(user.TrackingReference, 5000, "test transfer")
			if err != nil {
				t.Errorf("Error calling FundVirtualAccount: %s", err)
			}

			// Store the modified struct value back in the map
			users[i] = user
		} else {

			// Modify the field directly
			user.AccountNumber = ""

			// Store the modified struct value back in the map
			users[i] = user
		}

	}

	//
	userA := users[0]
	userB := users[1]

	// Call the VirtualFundTransfer function
	trackingRef := userA.TrackingReference
	beneficiaryAccountNumber := userB.AccountNumber
	beneficiaryBankCode := "999057"
	beneficiaryName := userB.FirstName + " " + userB.LastName
	senderName := userA.FirstName + " " + userA.LastName
	nameEnquirySessionID := ""
	amount := 500
	narration := "Test Transfer"
	response, err := k.VirtualFundTransfer(trackingRef, beneficiaryAccountNumber, beneficiaryBankCode, beneficiaryName, senderName, nameEnquirySessionID, amount, narration)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling VirtualFundTransfer: %s", err.Error())
		return
	}

	//
	if strings.Contains(response["message"].(string), "invalid") {
		// handle that
		t.Errorf("Error calling SingleFundTransfer: %s", response)
	}
}

func TestMainTxnLogs(t *testing.T) {
	// Initialize API client with test server's URL
	k := SetUpKudaBankAPI(params)

	// Call the MainTxnLogs function
	response, err := k.MainTxnLogs(10, 1)

	// Check for errors
	if err != nil {
		t.Errorf("Error calling MainTxnLogs: %s", err.Error())
		return
	}

	// Check response data
	status := response["status"].(bool)
	assert.Equal(t, true, status)

	// if response["data"].(map[string]interface{})["responseCode"] != "00" {
	// 	t.Errorf("Unexpected response code: %s", response["data"].(map[string]interface{})["responseCode"])
	// }
}

func TestFilterMainTxnLogs(t *testing.T) {
	k := SetUpKudaBankAPI(params) // initialize the API with your API key

	// call the API function with some test parameters
	response, err := k.FilterMainTxnLogs("2022-01-01T00:00:00.000Z", "2022-02-01T00:00:00.000Z", 10, 1)

	if err != nil {
		t.Errorf("FilterMainTxnLogs returned an error: %v", err)
	}

	// check that the response contains expected data
	// if _, ok := response["transactions"]; !ok {
	// 	t.Errorf("FilterMainTxnLogs did not return 'transactions' in response")
	// }
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestVirtualTxnLogs(t *testing.T) {
	k := SetUpKudaBankAPI(params)

	response, err := k.VirtualTxnLogs("3849991107", 10, 1)

	if err != nil {
		t.Errorf("VirtualTxnLogs returned an error: %v", err)
	}

	// if _, ok := response["transactions"]; !ok {
	// 	t.Errorf("VirtualTxnLogs did not return 'transactions' in response")
	// }
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestFilterVirtualTxnLogs(t *testing.T) {
	k := SetUpKudaBankAPI(params)

	response, err := k.FilterVirtualTxnLogs("3849991107", "2022-01-01T00:00:00.000Z", "2022-02-01T00:00:00.000Z", 10, 1)
	if err != nil {
		t.Errorf("FilterVirtualTxnLogs returned an error: %v", err)
	}

	// if _, ok := response["transactions"]; !ok {
	// 	t.Errorf("FilterVirtualTxnLogs did not return 'transactions' in response")
	// }
	status := response["status"].(bool)
	assert.Equal(t, true, status)
}

func TestFundVirtualAccount(t *testing.T) {
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(1)
	userA := users[0]

	// Call CreateVirtualAccount method
	userA.TrackingReference = strconv.Itoa(k.GetRef(10))
	response, err := k.CreateVirtualAccount(userA.FirstName, userA.LastName, userA.MiddleName, userA.Email, userA.PhoneNumber, userA.BusinessName, userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
	}
	status := response["status"].(bool)
	assert.Equal(t, true, status)

	response, err = k.FundVirtualAccount(userA.TrackingReference, 100000, "Test deposit")
	if err != nil {
		t.Errorf("FundVirtualAccount returned an error: %v", err)
	}
	status = response["status"].(bool)
	assert.Equal(t, true, status)
	// if _, ok := response["transactionId"]; !ok {
	// 	t.Errorf("FundVirtualAccount did not return 'transactionId' in response")
	// }

}

func TestWithdrawVirtualAccount(t *testing.T) {
	k := SetUpKudaBankAPI(params)
	users := CreateUsers(1)
	userA := users[0]

	// Call CreateVirtualAccount method
	userA.TrackingReference = strconv.Itoa(k.GetRef(10))
	response, err := k.CreateVirtualAccount(userA.FirstName, userA.LastName, userA.MiddleName, userA.Email, userA.PhoneNumber, userA.BusinessName, userA.TrackingReference)
	// Check for errors
	if err != nil {
		t.Errorf("Error calling CreateVirtualAccount: %s", err.Error())
	}
	status := response["status"].(bool)
	assert.Equal(t, true, status)

	// Fund Account Virtual Account
	response, err = k.FundVirtualAccount(userA.TrackingReference, 100000, "Test deposit")
	if err != nil {
		t.Errorf("FundVirtualAccount returned an error: %v", err)
	}
	status = response["status"].(bool)
	assert.Equal(t, true, status)

	// Withdraw From Virtual Account
	response, err = k.WithdrawVirtualAccount(userA.TrackingReference, 50000, "Test withdrawal")
	if err != nil {
		t.Errorf("WithdrawVirtualAccount returned an error: %v", err)
	}
	status = response["status"].(bool)
	assert.Equal(t, true, status)

	// if _, ok := response["transactionId"]; !ok {
	// 	t.Errorf("WithdrawVirtualAccount did not return 'transactionId' in response")
	// }

}

// govulncheck ./...
// TransactionStatusQuery()
// response, err := api.TransactionStatusQuery(request.IsThirdPartyBankTransfer, request.TransactionRequestReference)
