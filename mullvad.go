package mullvad

import "fmt"

type Client struct {
	// AccountNumber is the mullvad account number to use.
	AccountNumber string
	// Countries is a list of countries the configuration can be made from.
	Countries []string
}

func NewClient(accountNumber string, countryList []string) *Client {
	fmt.Println("NEW")
	return &Client{
		AccountNumber: accountNumber,
		Countries:     countryList,
	}
}

// will be our main client struct and constructor
