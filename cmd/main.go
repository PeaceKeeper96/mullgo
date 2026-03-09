package main

import (
	"fmt"
	"mullgo"
)

type Client struct {
	// AccountNumber is the mullvad account number to use.
	AccountNumber string
	// Countries is a list of countries the configuration can be made from.
	Countries []string
}

type Response struct {
	data string
}

func main() {
	sampleClient := mullvad.Client{
		AccountNumber: "", 
		Countries: []string{"Sweden"},
		}
	sampleRelay := &mullvad.Relay{
		Name: "testrelay",
		IPv4Addr: "123.123.123.123",
		MultihopPort: 45,
		PublicKey: "testkey",
		CityCode: "SAMP",
		Hostname: "hostnamesample",
	}
	
	config1, err := mullvad.GenerateConfig(
		sampleClient,
		[]mullvad.Relay{
			*sampleRelay,
		},
	)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(config1)

	jsonResp := Response{}
	mullvad.GetJson("https://api.mullvad.net/wg", &jsonResp)

	token := mullvad.GetToken(sampleClient.AccountNumber)
	fmt.Println("Token: ", token.AccessToken)
	fmt.Println("Expiry: ", token.Expiry)
}

// will be our main client struct and constructor
