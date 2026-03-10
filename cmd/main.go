package main

import (
	"fmt"
	mullvad "mullgo"
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
	sampleClient := mullvad.NewClient(
		"",
		[]string{"ch-zrh", "de-fra"},
	)
	// sampleRelay := &mullvad.Relay{
	// 	Name: "testrelay",
	// 	IPv4Addr: "123.123.123.123",
	// 	MultihopPort: 45,
	// 	PublicKey: "testkey",
	// 	CityCode: "SAMP",
	// 	Hostname: "hostnamesample",
	// }

	// config1, err := mullvad.GenerateConfig(
	// 	sampleClient,
	// 	[]mullvad.Relay{
	// 		*sampleRelay,
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println("ERROR")
	// }
	// fmt.Println(config1)

	// jsonResp := Response{}
	// mullvad.GetJson("https://api.mullvad.net/wg", &jsonResp)

	token := mullvad.GetToken(sampleClient.AccountNumber)
	fmt.Println("Token: ", token.AccessToken)
	fmt.Println("Expiry: ", token.Expiry)

	relays := mullvad.GetRelays(token)
	filteredRelays := mullvad.RelayFilter(*sampleClient, relays)
	fmt.Println("so back")
	fmt.Println(filteredRelays)

	pubKey, privKey := mullvad.Keygen()
	fmt.Println("PubKey: ", pubKey)
	fmt.Println("PrivKey: ", privKey)

	num, err := mullvad.AddKey(token, pubKey)
	fmt.Println(num, err)

	devices := mullvad.GetDevices(token)
	fmt.Println("Devices: ", devices)

	device := mullvad.GetDevicebyPubkey(devices, "CN7D6kEsXOlVSGFY5qAqym8ITCjA4NzGGERej8kxE3Q=")
	fmt.Println("\n", device)
}

// will be our main client struct and constructor
