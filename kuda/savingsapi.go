package kuda

// func (k *KudaBankAPI) GetSpendSavingsTxns(trackingRef int, pageNumber int, pageSize int) (JsonType, error) {
// 	serviceType := "RETRIEVE_SPEND_AND_SAVE_TRANSACTIONS"
// 	// make request and return response
// 	data := JsonType{}
// 	// Parse the response data and return it
// 	return data, nil

// }

// func (k *KudaBankAPI) CreatePlainSavings(savingsTrackingRef int, name string, virtualAccTrackingRef int) (JsonType, error) {
// 	serviceType := "CREATE_PLAIN_SAVINGS"
// 	// make request and return response
// 	data := JsonType{}
// 	// Parse the response data and return it
// 	return data, nil
// }

// func (k *KudaBankAPI) GetPlainSavings(trackingRef int) (JsonType, error) {
// 	serviceType := "GET_PLAIN_SAVINGS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllCustomerPlainSavings(accountNumber int) (JsonType, error) {
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllPlainSavings() (JsonType, error) {
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) UpdatePlainSavingsAccount(trackingRef int, status int) (JsonType, error) {
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) PlainSavingsTxn(trackingRef int, amount int, transactionType string, narration string) (JsonType, error) {
// 	serviceType := "PLAIN_SAVINGS_TRANSACTIONS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) PlainSavingsTxnLogs(trackingRef int, pageSize int, pageNumber int) (JsonType, error) {
// 	serviceType := "RETRIEVE_PLAIN_SAVINGS_TRANSACTIONS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) CreateOpenFlexibleSavings(savingsTrackingRef string, name string, virtualAccTrackingRef string, amount int, duration int, frequency int, startNow string, startDate string, isInterestEarning bool) (JsonType, error) {
// 	serviceType := "CREATE_OPEN_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) CreateClosedFlexibleSave(savingsTrackingRef string, name string, virtualAccTrackingRef string, amount int, duration int, frequency int, startDate string, startNow bool, isInterestEarning bool) (JsonType, error) {
// 	serviceType := "CREATE_CLOSED_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetOpenFlexibleSave(trackingRef int) (JsonType, error) {
// 	serviceType := "GET_OPEN_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetClosedFlexibleSavings(trackingRef int) (JsonType, error) {
// 	serviceType := "GET_CLOSED_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllCustomerOpenFlexibleSavings(primaryAccountNumber int) (JsonType, error) {
// 	serviceType := "GET_ALL_CUSTOMER_OPEN_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllCustomerClosedFlexibleSavings(primaryAccountNumber int) (JsonType, error) {
// 	serviceType := "GET_ALL_CUSTOMER_CLOSED_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllOpenFlexibleSavings() (JsonType, error) {
// 	serviceType := "GET_ALL_OPEN_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllClosedFlexibleSavings() (JsonType, error) {
// 	serviceType := "GET_ALL_CLOSED_FLEXIBLE_SAVE"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) WithdrawOpenFlexibleSavings() (JsonType, error) {
// 	serviceType := "COMPLETE_OPEN_FLEXIBLE_SAVE_WITHDRAWAL"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) WithdrawClosedFlexibleSavings() (JsonType, error) {
// 	serviceType := "COMPLETE_CLOSED_FLEXIBLE_SAVE_WITHDRAWAL"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) OpenFlexibleSavingsTxnLogs(trackingRef int, pageSize int, pageNumber int) (JsonType, error) {
// 	serviceType := "RETRIEVE_OPEN_FLEXIBLE_SAVINGS_TRANSACTIONS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) ClosedFlexibleSavingsTxnLogs(trackingRef int, pageSize int, pageNumber int) (JsonType, error) {
// 	serviceType := "RETRIEVE_CLOSED_FLEXIBLE_SAVINGS_TRANSACTIONS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) CreateFixedSavings(savingsTrackingRef int, name string, virtualAccTrackingRef int, amount int, duration int, startDate string, startNow bool, isInterestEarning bool) (JsonType, error) {
// 	serviceType := "CREATE_FIXED_SAVINGS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetFixedSavings(trackingRef int) (JsonType, error) {
// 	serviceType := "GET_FIXED_SAVINGS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllCustomerFixedSavings(primaryAccountNumber int) (JsonType, error) {
// 	serviceType := "GET_ALL_CUSTOMER_FIXED_SAVINGS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) GetAllFixedSavings() (JsonType, error) {
// 	serviceType := "GET_ALL_FIXED_SAVINGS"
// 	return make(JsonType), nil
// }

// func (k *KudaBankAPI) TerminateFixedDeposit(trackingRef int, amount int) (JsonType, error) {

// 	payload := JsonType{
// 		"trackingReference": trackingRef,
// 		"amount":            amount,
// 	}

// 	resp, err := k.MakeRequest(COMPLETE_FIXED_SAVINGS_WITHDRAWAL, payload, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

// func (k *KudaBankAPI) FixedSavingsTxnLogs(trackingRef int, pageSize int, pageNumber int) (JsonType, error) {

// 	payload := JsonType{
// 		"trackingReference": trackingRef,
// 		"pageSize":          pageSize,
// 		"pageNumber":        pageNumber,
// 	}

// 	resp, err := k.MakeRequest(RETRIEVE_FIXED_SAVINGS_TRANSACTIONS, payload, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
