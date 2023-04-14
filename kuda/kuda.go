package kuda

import (
	"bytes"
	"math/rand"

	// "crypto/rand"

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// NewKudaBankAPI returns a new instance of KudaBankAPI.
func NewKudaBankAPI(connectionParams *ConnectionParams) (*KudaBankAPI, error) {
	// Validate input values
	if connectionParams.Email == "" {
		return nil, errors.New("email is required")
	}
	if connectionParams.ApiKey == "" {
		return nil, errors.New("api key is required")
	}

	if connectionParams.BaseURL == "" {
		if !connectionParams.Live {
			connectionParams.BaseURL = "https://kuda-openapi-uat.kudabank.com/v2.1"
		} else {
			connectionParams.BaseURL = "https://kuda-openapi.kuda.com/v2.1"
		}
	}

	return &KudaBankAPI{
		client:           &http.Client{},
		connectionParams: connectionParams,
	}, nil
}

func (k *KudaBankAPI) GetToken() (string, error) {

	url := k.connectionParams.BaseURL + "/Account/GetToken"

	requestData := map[string]string{
		"email":  k.connectionParams.Email,
		"apikey": k.connectionParams.ApiKey,
	}

	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	resp, err := k.client.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// fmt.Println(string(bodyBytes))
	response := string(bodyBytes)

	return response, nil
}

func (k *KudaBankAPI) MakeRequest(serviceType ServiceType, payload JsonType, requestRef string) (JsonType, error) {

	token, err := k.GetToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", k.connectionParams.BaseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	if requestRef == "" {
		// randomBytes := make([]byte, 10)
		// rand.Read(randomBytes)
		// hexRef := hex.EncodeToString(randomBytes)
		// hexRef := strconv.Itoa(k.GetRef(10))
		// requestRef = &hexRef
		requestRef = strconv.Itoa(k.GetRef(10))
	}

	requestData := JsonType{
		"serviceType": string(serviceType),
		"requestRef":  requestRef,
		"data":        payload,
	}

	// jsonBytes, err := json.Marshal(JsonType{"data": requestData})
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(jsonBytes))
	req.ContentLength = int64(len(jsonBytes))

	resp, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(bodyBytes))

	var response JsonType
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		response = JsonType{
			"Status":  false,
			"Message": json.Unmarshal(bodyBytes, &JsonType{}),
		}
	}

	return response, nil
}

// GetRef generates a unique reference per request (requestRef)
func (k *KudaBankAPI) GetRef(length int) int {
	const digits = "0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(digits[rand.Intn(len(digits))])
	}
	resultInt, _ := strconv.Atoi(sb.String())
	return resultInt
}

// KoboToNGN converts a kobo amount to naira.
func (k *KudaBankAPI) KoboToNGN(kobo int) float64 {
	return float64(kobo) / 100
}

// NGNToKobo converts a naira amount to kobo.
func (k *KudaBankAPI) NGNToKobo(naira float64) int {
	return int(naira * 100)
}
