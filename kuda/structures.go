package kuda

import "net/http"

// Credentials
type ConnectionParams struct {
	Email       string
	ApiKey      string
	Live        bool
	ShowRequest bool
	ShowHeader  bool
	BaseURL     string
}

// KudaBankAPI provides methods for interacting with the Kuda API.
type KudaBankAPI struct {
	client           *http.Client
	connectionParams *ConnectionParams
}

// for the json data structure
type JsonType map[string]interface{}

// Service types
type ServiceType string
