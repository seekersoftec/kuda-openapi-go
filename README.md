# Kuda OpenAPI

This package contains a Go program that uses the go-faker and kuda-openapi-go packages to create personas, virtual accounts and perform single fund transfers using the Kuda bank API.

## Installation

To use this package, you need to have Go installed on your computer.

1. Clone the repository to your local machine.
2. Install the dependencies:

```bash
    go get github.com/go-faker/faker/v4
    go get github.com/seekersoftec/kuda-openapi-go
```

3. Navigate to the cloned directory and run the main.go file.

## Usage

The program generates fake data using the faker package and creates virtual accounts and performs single fund transfers using the Kuda bank API. The generateFakeData() function generates fake data that is used to create virtual accounts. The normalizePhoneNumber() function normalizes phone numbers by replacing the plus sign with zero and replacing the first three digits with a random value. The main() function creates two personas, creates a virtual account for each persona, and then performs a single fund transfer from one virtual account to another.

## Required parameters

The following parameters are required to use the program:

- Email: your email address
- ApiKey: your API key
- Live: false (if you want to test on a sandbox environment)
- ShowRequest: false
- ShowHeader: false

## License

This package is licensed under the MIT License.
