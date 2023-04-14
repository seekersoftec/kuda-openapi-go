package kuda

func (k *KudaBankAPI) NameEnquiry(accountNumber string, bankCode string, senderTrackingReference string, requestRef string) (JsonType, error) {
	/*Validate and retrieve information on a NUBAN account number
	  https://kudabank.gitbook.io/kudabank/single-fund-transfer/name-enquiry
	  Args:
	      accountNumber (int): Destination Account Number
	      bankCode (int): Destination Bank Code. Defaults to Kuda bank code
		  senderTrackingReference (string): senderTrackingReference
		  requestRef (string): requestRef
	*/
	isRequestFromVirtualAccount := "true"

	if senderTrackingReference == "" {
		senderTrackingReference = ""
		isRequestFromVirtualAccount = "false"
	}

	payload := JsonType{
		"beneficiaryAccountNumber":    accountNumber,
		"beneficiaryBankCode":         bankCode,
		"SenderTrackingReference":     senderTrackingReference,
		"isRequestFromVirtualAccount": isRequestFromVirtualAccount,
	}

	return k.MakeRequest(NAME_ENQUIRY, payload, requestRef)
}

func (k *KudaBankAPI) BankList() (JsonType, error) {
	/*The Bank list covered here are for all Nigerian Financial Institutions supported by NIBSS
	  https://kudabank.gitbook.io/kudabank/single-fund-transfer/bank-list

	*/
	return k.MakeRequest(BANK_LIST, JsonType{}, "")
}

func (k *KudaBankAPI) CreateVirtualAccount(firstName string, lastName string, middleName string, email string, phoneNumber string, businessName string, trackingRef string) (JsonType, error) {
	/*Create a Virtual Account
	  https://kudabank.gitbook.io/kudabank/virtual-account-creation#create-a-virtual-account
	  Args:
	      lastName (str): User's last name
	      firstName (str): User's first name
	      emailAddress (str): User's email address
	      phoneNumber (str): User's phone number
	      trackingRef (int): Unique identifier for the account

	*/

	payload := JsonType{
		"email":             email,
		"phoneNumber":       phoneNumber,
		"lastName":          lastName,
		"firstName":         firstName,
		"middleName":        middleName,
		"businessName":      businessName,
		"trackingReference": trackingRef,
	}

	return k.MakeRequest(ADMIN_CREATE_VIRTUAL_ACCOUNT, payload, "")
}

func (k *KudaBankAPI) GetSingleVirtualAccount(trackingRef string) (JsonType, error) {
	/*Retrieve Virtual Account information
	  https://kudabank.gitbook.io/kudabank/virtual-account-creation/retrieve-virtual-account
	  Args:
	      trackingRef (int): Unique Identifier for virtual Account
	  Returns:
	      JsonType
	*/
	payload := JsonType{
		"trackingReference": trackingRef,
	}
	return k.MakeRequest(RETRIEVE_SINGLE_VIRTUAL_ACCOUNT, payload, "")
}

func (k *KudaBankAPI) GetVirtualAccountBalance(trackingRef string) (JsonType, error) {
	/*
		Get Virtual Account Balance

		  Args:
		      trackingRef (int): Unique Identifier for virtual Account
		  Returns:
		      JsonType
	*/
	payload := JsonType{
		"trackingReference": trackingRef,
	}
	return k.MakeRequest(RETRIEVE_VIRTUAL_ACCOUNT_BALANCE, payload, "")
}

func (k *KudaBankAPI) SingleFundTransfer(trackingRef string, beneficiaryAccountNumber string, beneficiaryBankCode string, beneficiaryName string, amount int, senderName string, narration string, ClientAccountNumber string, nameEnquirySessionID string) (JsonType, error) {
	/*Bank transfer from a Kuda Account to any bank account

	  https://kudabank.gitbook.io/kudabank/single-fund-transfer/send-money-from-a-kuda-account
	  Args:
	      trackingRef (int): Request Reference ID
	      beneficiaryAccountNumber (int): Destination bank account number
	      beneficiaryBankCode (int): Destination bank code
	      beneficiaryName (str): Name of
	*/
	payload := JsonType{
		"ClientAccountNumber":  ClientAccountNumber,
		"beneficiarybankCode":  beneficiaryBankCode,
		"beneficiaryAccount":   beneficiaryAccountNumber,
		"beneficiaryName":      beneficiaryName,
		"amount":               amount,
		"narration":            narration,
		"nameEnquirySessionID": nameEnquirySessionID,
		"trackingReference":    trackingRef,
		"senderName":           senderName,
		"clientFeeCharge":      0,
	}

	if nameEnquirySessionID == "" {
		nameEnquiryResult, err := k.NameEnquiry(beneficiaryAccountNumber, beneficiaryBankCode, "", "")
		if err != nil {
			return nil, err
		}
		// fmt.Println(nameEnquiryResult)
		// tempData := nameEnquiryResult["data"].(map[string]interface{})
		// fmt.Println(tempData)

		payload["beneficiaryName"] = nameEnquiryResult["beneficiaryName"]
		payload["nameEnquirySessionID"] = nameEnquiryResult["sessionID"]
	}

	// fmt.Println("SINGLE_FUND_TRANSFER", payload)
	return k.MakeRequest(SINGLE_FUND_TRANSFER, payload, "")
}

func (k *KudaBankAPI) VirtualFundTransfer(trackingRef string, beneficiaryAccountNumber string, beneficiaryBankCode string, beneficiaryName string, senderName string, nameEnquirySessionID string, amount int, narration string) (JsonType, error) {
	/*
			Transfer money from a virtual account
			Bank transfer from a KUDA Virtual Account to any bank account
		    https://kudabank.gitbook.io/kudabank/single-fund-transfer/virtual-account-fund-transfer
		    Args:
		        trackingRef (int): Unique Identifier for virtual Account
		        beneficiaryAccount (int): Destination bank account number
		        beneficiaryBankCode (int): Destination bank code
		        beneficiaryName (str): Name of the recipient
		        amount (int): Amount to be transferred (in kobo)
		        narration (str): User defined reason for the transaction.
		        SenderName (str): Name of the person sending money
		        nameEnquirySessionID (str, optional): Session ID generated from the nameEnquiry request. Defaults to None.

	*/

	payload := JsonType{
		"trackingReference":   trackingRef,
		"beneficiarybankCode": beneficiaryBankCode,
		"beneficiaryAccount":  beneficiaryAccountNumber,
		"beneficiaryName":     beneficiaryName,
		"amount":              amount,
		"narration":           narration,
		"SenderName":          senderName,
		"nameEnquiryId":       nameEnquirySessionID,
		"clientFeeCharge":     0,
	}

	// TODO: Check Request Body
	// var accountName string
	if nameEnquirySessionID == "" {
		nameEnquiryResult, err := k.NameEnquiry(beneficiaryAccountNumber, beneficiaryBankCode, "", "")
		if err != nil {
			return nil, err
		}
		// fmt.Println(nameEnquiryResult)
		// tempData := nameEnquiryResult["data"].(map[string]interface{})
		// fmt.Println(tempData)

		payload["beneficiaryName"] = nameEnquiryResult["beneficiaryName"]
		payload["nameEnquirySessionID"] = nameEnquiryResult["sessionID"]
	}

	// fmt.Println("VIRTUAL_ACCOUNT_FUND_TRANSFER payload => ", payload)
	return k.MakeRequest(VIRTUAL_ACCOUNT_FUND_TRANSFER, payload, "")
}

func (k *KudaBankAPI) MainTxnLogs(pageSize int, pageNumber int) (JsonType, error) {
	/*
			Retrieve a list of all transactions for the currently authenticated user.
		    https://kudabank.gitbook.io/kudabank/view-transaction-history/kuda-account-transaction-history
		    Args:
		        pageSize (int)
		        pageNumber (int)
	*/

	payload := JsonType{
		"pageSize":   pageSize,
		"pageNumber": pageNumber,
	}

	return k.MakeRequest(ADMIN_MAIN_ACCOUNT_TRANSACTIONS, payload, "")
}

func (k *KudaBankAPI) FilterMainTxnLogs(startDate string, endDate string, pageSize int, pageNumber int) (JsonType, error) {
	/*
			Retrieve a filtered list of all transactions for the currently authenticated user.
		    https://kudabank.gitbook.io/kudabank/view-transaction-history/filtered-kuda-account-transaction-history
		    Args:
		        startDate (str): Ex: 2020-10-27T09:58:23.4740446Z
		        endDate (str): Ex: 2020-12-27T09:58:23.4740446Z
		        pageSize (int, optional)
		        pageNumber (int, optional)
	*/

	payload := JsonType{
		"pageSize":   pageSize,
		"pageNumber": pageNumber,
		"startDate":  startDate,
		"endDate":    endDate,
	}

	return k.MakeRequest(ADMIN_MAIN_ACCOUNT_FILTERED_TRANSACTIONS, payload, "")
}

func (k *KudaBankAPI) VirtualTxnLogs(trackingRef string, pageSize int, pageNumber int) (JsonType, error) {
	/*
			Retrieve a list of all transactions for a specified virtual account
		    https://kudabank.gitbook.io/kudabank/view-transaction-history/virtual-accounttransactionhistory
		    Args:
		        trackingRef (string): Unique Identifier for virtual Account
		        pageSize (int, optional)
		        pageNumber (int, optional)
	*/
	payload := JsonType{
		"trackingReference": trackingRef,
		"pageSize":          pageSize,
		"pageNumber":        pageNumber,
	}

	return k.MakeRequest(ADMIN_VIRTUAL_ACCOUNT_TRANSACTIONS, payload, "")
}

func (k *KudaBankAPI) FilterVirtualTxnLogs(trackingRef string, startDate string, endDate string, pageSize int, pageNumber int) (JsonType, error) {
	/*
			Retrieve a filtered list of all transactions for a specified virtual account
		    https://kudabank.gitbook.io/kudabank/view-transaction-history/virtual-account-transaction-history
		    Args:
		        trackingRef (string): Unique Identifier for virtual Account
		        startDate (str): Ex: 2020-10-27T09:58:23.4740446Z
		        endDate (str): Ex: 2020-12-27T09:58:23.4740446Z
		        pageSize (int, optional)
		        pageNumber (int, optional)
	*/

	payload := JsonType{
		"trackingReference": trackingRef,
		"pageSize":          pageSize,
		"pageNumber":        pageNumber,
		"startDate":         startDate,
		"endDate":           endDate,
	}

	return k.MakeRequest(ADMIN_MAIN_ACCOUNT_FILTERED_TRANSACTIONS, payload, "")
}

func (k *KudaBankAPI) FundVirtualAccount(trackingRef string, amount int, narration string) (JsonType, error) {
	/*
			Deposit to a virtual account
		    https://kudabank.gitbook.io/kudabank/add-remove-money-from-a-virtual-account
		    Args:
		        trackingRef (string): Unique Identifier for virtual Account
		        amount (int): Amount to fund (in kobo)
		        narration (str, optional): Transaction Description. Defaults to "".
	*/
	payload := JsonType{
		"trackingReference": trackingRef,
		"amount":            amount,
		"narration":         narration,
	}

	return k.MakeRequest(FUND_VIRTUAL_ACCOUNT, payload, "")
}

func (k *KudaBankAPI) WithdrawVirtualAccount(trackingRef string, amount int, narration string) (JsonType, error) {
	/*
			Withdrawing funds from a virtual account means to transfer funds from a virtual account to an associated KUDA account or to any other Nigerian Bank account.
		    https://kudabank.gitbook.io/kudabank/add-remove-money-from-a-virtual-account#withdraw-from-virtual-account
		    Args:
		        trackingRef (string): Unique Identifier for virtual Account
		        amount (int): Amount to be withdrawn (in kobo)
		        narration (str, optional): Transaction description. Defaults to "".
	*/

	payload := JsonType{
		"trackingReference": trackingRef,
		"amount":            amount,
		"narration":         narration,
	}

	return k.MakeRequest(WITHDRAW_VIRTUAL_ACCOUNT, payload, "")
}
