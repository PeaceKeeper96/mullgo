# mullgo
A go package for using the mullvad API and generating wireguard keys.

The package does not generate configurations for you, but gives you everything necessary to do this. You can also **delete** devices from mullvad as long as you know either the ID, Name or pubkey.

# Example Usage:
```go
package main

import (
	"fmt"
	mullvad "github.com/PeaceKeeper96/mullgo"
)

func main() {
	sampleClient := mullvad.NewClient(
		"123456789", // Account Number goes here
		[]string{"ch-zrh", "de-fra"}, // List of locations you want to use
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
```

# List of locations

```go
"al-tia", // Tirana, Albania
"ar-bue", // Buenos Aires, Argentina
"at-vie", // Vienna, Austria
"au-adl", // Adelaide, Australia
"au-bne", // Brisbane, Australia
"au-mel", // Melbourne, Australia
"au-per", // Perth, Australia
"au-syd", // Sydney, Australia
"be-bru", // Brussels, Belgium
"bg-sof", // Sofia, Bulgaria
"br-for", // Fortaleza, Brazil
"br-sao", // São Paulo, Brazil
"ca-mtr", // Montreal, Canada
"ca-tor", // Toronto, Canada
"ca-van", // Vancouver, Canada
"ca-yyc", // Calgary, Canada
"ch-zrh", // Zurich, Switzerland
"cl-scl", // Santiago, Chile
"co-bog", // Bogotá, Colombia
"cy-nic", // Nicosia, Cyprus
"cz-prg", // Prague, Czech Republic
"de-ber", // Berlin, Germany
"de-dus", // Düsseldorf, Germany
"de-fra", // Frankfurt, Germany
"dk-cph", // Copenhagen, Denmark
"ee-tll", // Tallinn, Estonia
"es-bcn", // Barcelona, Spain
"es-mad", // Madrid, Spain
"es-vlc", // Valencia, Spain
"fi-hel", // Helsinki, Finland
"fr-bod", // Bordeaux, France
"fr-mrs", // Marseille, France
"fr-par", // Paris, France
"gb-glw", // Glasgow, United Kingdom
"gb-lon", // London, United Kingdom
"gb-mnc", // Manchester, United Kingdom
"gr-ath", // Athens, Greece
"hk-hkg", // Hong Kong
"hr-zag", // Zagreb, Croatia
"hu-bud", // Budapest, Hungary
"id-jpu", // Jakarta, Indonesia
"ie-dub", // Dublin, Ireland
"il-tlv", // Tel Aviv, Israel
"it-mil", // Milan, Italy
"it-pmo", // Palermo, Italy
"jp-osa", // Osaka, Japan
"jp-tyo", // Tokyo, Japan
"mx-qro", // Querétaro, Mexico
"my-kul", // Kuala Lumpur, Malaysia
"ng-los", // Lagos, Nigeria
"nl-ams", // Amsterdam, Netherlands
"no-osl", // Oslo, Norway
"no-svg", // Stavanger, Norway
"nz-akl", // Auckland, New Zealand
"pe-lim", // Lima, Peru
"ph-mnl", // Manila, Philippines
"pl-waw", // Warsaw, Poland
"pt-lis", // Lisbon, Portugal
"ro-buh", // Bucharest, Romania
"rs-beg", // Belgrade, Serbia
"se-got", // Gothenburg, Sweden
"se-mma", // Malmö, Sweden
"se-sto", // Stockholm, Sweden
"sg-sin", // Singapore
"si-lju", // Ljubljana, Slovenia
"sk-bts", // Bratislava, Slovakia
"th-bkk", // Bangkok, Thailand
"tr-ist", // Istanbul, Turkey
"ua-iev", // Kyiv, Ukraine
"us-atl", // Atlanta, USA
"us-bos", // Boston, USA
"us-chi", // Chicago, USA
"us-dal", // Dallas, USA
"us-den", // Denver, USA
"us-det", // Detroit, USA
"us-hou", // Houston, USA
"us-lax", // Los Angeles, USA
"us-mia", // Miami, USA
"us-mkc", // Kansas City, USA
"us-nyc", // New York City, USA
"us-phx", // Phoenix, USA
"us-qas", // Quincy, USA
"us-rag", // Raleigh, USA
"us-sea", // Seattle, USA
"us-sjc", // San Jose, USA
"us-slc", // Salt Lake City, USA
"us-txc", // Texas City, USA
"us-uyk", // Unknown / internal code (US)
"us-was", // Washington D.C., USA
"za-jnb", // Johannesburg, South Africa
```