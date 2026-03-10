package mullvad

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Access token information, including expiry.
type Token struct {
	AccessToken string    `json:"access_token"`
	Expiry      time.Time `json:"expiry"`
}

// A list of all mullvad devices and their information
type Devices []Device

// Individual mullvad device information
type Device struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Pubkey      string        `json:"pubkey"`
	HijackDNS   bool          `json:"hijack_dns"`
	Created     time.Time     `json:"created"`
	Ipv4Address string        `json:"ipv4_address"`
	Ipv6Address string        `json:"ipv6_address"`
	Ports       []interface{} `json:"ports"`
}

type Account struct {
	ID            string    `json:"id"`
	Expiry        time.Time `json:"expiry"`
	HasPayments   bool      `json:"has_payments"`
	MaxPorts      int       `json:"max_ports"`
	CanAddPorts   bool      `json:"can_add_ports"`
	MaxDevices    int       `json:"max_devices"`
	CanAddDevices bool      `json:"can_add_devices"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// Get authorization token from mullvad api
func GetToken(accountNumber string) (Token, error) {
	if accountNumber == "" {
		return Token{}, fmt.Errorf("GetToken: No Account number supplied.\n")
	}
	client := myClient

	url := "https://api.mullvad.net/auth/v1/token"
	method := "POST"
	postdata := strings.NewReader(`{"account_number":"` + accountNumber + `"}`)
	
	req, err := http.NewRequest(method, url, postdata)
	if err != nil {
        return Token{}, fmt.Errorf("GetToken: failed to create request: %w", err)
	}
	// Set header
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
        return Token{}, fmt.Errorf("GetToken: failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read and print response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        return Token{}, fmt.Errorf("GetToken: failed to read response: %w", err)
	}

	var unmarshaledToken Token
	if err := json.Unmarshal(body, &unmarshaledToken); err != nil {
        return Token{}, fmt.Errorf("GetToken: failed to parse response: %w", err)
	}

	return unmarshaledToken, nil
}

// get up to date relay information from mullvad
func GetRelays(token Token) (RelayList, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := "https://api.mullvad.net/app/v1/relays"
	getdata := strings.NewReader(`{"access_token":"` + token.AccessToken + `"}`)

	req, err := http.NewRequest("GET", url, getdata)
	if err != nil {
        return RelayList{}, fmt.Errorf("GetRelays: failed to create request: %w", err)
	}
	// Set header
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
        return RelayList{}, fmt.Errorf("GetRelays: failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read and print response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        return RelayList{}, fmt.Errorf("GetRelays: failed to read response: %w", err)
	}

	var unmarshaledRelays RelayList
	if err := json.Unmarshal(body, &unmarshaledRelays); err != nil {
        return RelayList{}, fmt.Errorf("GetRelays: failed to parse response: %w", err)
	}

	return unmarshaledRelays, nil
}

// Add a wireguard key to the mullvad account
func AddKey(token Token, pubKey string) (int, error) {
	client := myClient

	formData := url.Values{}
	formData.Set("key", pubKey)

	req, err := http.NewRequest("POST", "https://mullvad.net/en/account/devices?/upload-key", strings.NewReader(formData.Encode()))
	if err != nil {
		return 0, fmt.Errorf("AddKey: failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Referer", "https://mullvad.net/en/account/devices")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Add("Origin", "https://mullvad.net")
	req.Header.Add("x-sveltekit-action", "true")
	req.Header.Add("User-Agent", "Go")
	req.AddCookie(&http.Cookie{
		Name:  "accessToken",
		Value: token.AccessToken,
	})

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("AddKey: failed to make request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func GetDevices(token Token) (Devices, error) {
	var allDevices Devices

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://api.mullvad.net/accounts/v1/devices", nil)
	if err != nil {
		return Devices{}, fmt.Errorf("GetDevices: failed to create request: %w", err)
	}

	authtoken := "Bearer " + token.AccessToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")

	// Set header
	resp, err := client.Do(req)
	if err != nil {
		return Devices{}, fmt.Errorf("GetDevices: failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Devices{}, fmt.Errorf("GetDevices: failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &allDevices); err != nil {
		return Devices{}, fmt.Errorf("GetDevices: failed to parse response: %w", err)
	}

	return allDevices, nil
}

// Get a device by the public key associated to it
func GetDevicebyPubkey(deviceList Devices, pubkey string) (Device, error) {
	if pubkey == "" {
		return Device{}, fmt.Errorf("GetDevicebyPubkey: No pubkey was supplied")
	}
	for i := range deviceList {
		if deviceList[i].Pubkey == pubkey {
			return deviceList[i], nil
		}
	}
	return Device{}, fmt.Errorf("GetDevicebyPubkey: failed to find device")
}

// Get a device by the internal id used by the api
func GetDevicebyId(deviceList Devices, id string) (Device, error) {
	if id == "" {
		return Device{}, fmt.Errorf("GetDevicebyId: No id was supplied")
	}
	for i := range deviceList {
		if deviceList[i].ID == id {
			return deviceList[i], nil
		}
	}
	return Device{}, fmt.Errorf("GetDevicebyId: failed to find device")
}

// Get a device by searching the human readeable name e.g: agile boa
func GetDevicebyName(deviceList Devices, name string) (Device, error) {
	if name == "" {
		return Device{}, fmt.Errorf("GetDevicebyName: No name was supplied")
	}
	for i := range deviceList {
		if deviceList[i].Name == name {
			return deviceList[i], nil
		}
	}
	return Device{}, fmt.Errorf("GetDevicebyName: failed to find device")
}

func RemoveDevice(token Token, id string) (int, error) {
	client := myClient

	formData := url.Values{}
	formData.Set("id", id)

	req, err := http.NewRequest("POST", "https://mullvad.net/en/account/devices?/revoke-device", strings.NewReader(formData.Encode()))
	if err != nil {
		return 0, fmt.Errorf("RemoveDevice: failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Referer", "https://mullvad.net/en/account/devices")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Add("Origin", "https://mullvad.net")
	req.Header.Add("x-sveltekit-action", "true")
	req.Header.Add("User-Agent", "Go")
	req.AddCookie(&http.Cookie{
		Name:  "accessToken",
		Value: token.AccessToken,
	})

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("RemoveDevice: failed to make request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func GetAccountInfo(token Token) (Account, error) {
	var currentAccount Account

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://api.mullvad.net/accounts/v1/accounts/me", nil)
	if err != nil {
        return Account{}, fmt.Errorf("GetAccountInfo: failed to create request: %w", err)
	}

	authtoken := "Bearer " + token.AccessToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
        return Account{}, fmt.Errorf("GetAccountInfo: failed to make request: %w", err)	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
        return Account{}, fmt.Errorf("GetAccountInfo: failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &currentAccount); err != nil {
        return Account{}, fmt.Errorf("GetAccountInfo: failed to parse response: %w", err)
	}

	return currentAccount, nil
}