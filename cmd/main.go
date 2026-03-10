package main

import (
	"fmt"
	mullgo "github.com/PeaceKeeper96/mullgo"
)

func main() {
	sampleClient := mullgo.NewClient(
		"",
		[]string{"ch-zrh", "de-fra"},
	)

	token, _ := mullgo.GetToken(sampleClient.AccountNumber)
	fmt.Println("Token: ", token.AccessToken)
	fmt.Println("Expiry: ", token.Expiry)

	relays, _ := mullgo.GetRelays(token)
	filteredRelays := mullgo.RelayFilter(*sampleClient, relays)
	fmt.Println("so back")
	fmt.Println(filteredRelays)

	pubKey, privKey, _ := mullgo.Keygen()
	fmt.Println("PubKey: ", pubKey)
	fmt.Println("PrivKey: ", privKey)

	keyCode, err := mullgo.AddKey(token, pubKey)
	fmt.Println(keyCode, err)

	devices, _ := mullgo.GetDevices(token)
	fmt.Println("Devices: ", devices)

	device, _ := mullgo.GetDevicebyPubkey(devices, devices[0].Pubkey)
	fmt.Println("\n", device)

	removeCode, _ := mullgo.RemoveDevice(token, device.ID)
	fmt.Println("\nIf removed, you will see 200: ", removeCode)

	accountInfo, _ := mullgo.GetAccountInfo(token)
	fmt.Println("\nAccount Info: ", accountInfo)
}

// will be our main client struct and constructor
