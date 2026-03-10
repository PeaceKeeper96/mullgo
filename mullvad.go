package mullvad

type Client struct {
	// AccountNumber is the mullvad account number to use.
	AccountNumber string
	// Countries is a list of countries the configuration can be made from.
	Locations []string
}

// Create an individual client config
func NewClient(accountNumber string, locationList []string) *Client {
	return &Client{
		AccountNumber: accountNumber,
		Locations:     locationList,
	}
}
