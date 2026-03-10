package main

import (
	"fmt"
	mullvad "mullgo"
)

func main() {
	sampleClient := mullvad.NewClient(
		"",
		[]string{"ch-zrh", "de-fra"},
	)

	token, _ := mullvad.GetToken(sampleClient.AccountNumber)
	fmt.Println("Token: ", token.AccessToken)
	fmt.Println("Expiry: ", token.Expiry)

	relays, _ := mullvad.GetRelays(token)
	filteredRelays := mullvad.RelayFilter(*sampleClient, relays)
	fmt.Println("so back")
	fmt.Println(filteredRelays)

	pubKey, privKey, _ := mullvad.Keygen()
	fmt.Println("PubKey: ", pubKey)
	fmt.Println("PrivKey: ", privKey)

	keyCode, err := mullvad.AddKey(token, pubKey)
	fmt.Println(keyCode, err)

	devices, _ := mullvad.GetDevices(token)
	fmt.Println("Devices: ", devices)

	device, _ := mullvad.GetDevicebyPubkey(devices, devices[0].Pubkey)
	fmt.Println("\n", device)

	removeCode, _ := mullvad.RemoveDevice(token, device.ID)
	fmt.Println("\nIf removed, you will see 200: ", removeCode)

	accountInfo, _ := mullvad.GetAccountInfo(token)
	fmt.Println("\nAccount Info: ", accountInfo)
}

// will be our main client struct and constructor
