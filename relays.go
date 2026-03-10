package mullgo

import (
	"encoding/json"
	"slices"
)

type RelayList struct {
	Locations struct {
		AuAdl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"au-adl"`
		NlAms struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"nl-ams"`
		UsQas struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-qas"`
		GrAth struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"gr-ath"`
		UsAtl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-atl"`
		NzAkl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"nz-akl"`
		ThBkk struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"th-bkk"`
		EsBcn struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"es-bcn"`
		RsBeg struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"rs-beg"`
		DeBer struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"de-ber"`
		CoBog struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"co-bog"`
		FrBod struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"fr-bod"`
		UsBos struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-bos"`
		SkBts struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"sk-bts"`
		AuBne struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"au-bne"`
		BeBru struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"be-bru"`
		RoBuh struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ro-buh"`
		HuBud struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"hu-bud"`
		ArBue struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ar-bue"`
		CaYyc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ca-yyc"`
		UsChi struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-chi"`
		DkCph struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"dk-cph"`
		UsDal struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-dal"`
		UsDen struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-den"`
		UsDet struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-det"`
		IeDub struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ie-dub"`
		DeDus struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"de-dus"`
		BrFor struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"br-for"`
		DeFra struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"de-fra"`
		GbGlw struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"gb-glw"`
		SeGot struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"se-got"`
		FiHel struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"fi-hel"`
		HkHkg struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"hk-hkg"`
		UsHou struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-hou"`
		TrIst struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"tr-ist"`
		IDJpu struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"id-jpu"`
		ZaJnb struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"za-jnb"`
		UsMkc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-mkc"`
		MyKul struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"my-kul"`
		UaIev struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ua-iev"`
		NgLos struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ng-los"`
		PeLim struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"pe-lim"`
		PtLis struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"pt-lis"`
		SiLju struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"si-lju"`
		GbLon struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"gb-lon"`
		UsLax struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-lax"`
		EsMad struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"es-mad"`
		SeMma struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"se-mma"`
		GbMnc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"gb-mnc"`
		PhMnl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ph-mnl"`
		FrMrs struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"fr-mrs"`
		UsTxc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-txc"`
		AuMel struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"au-mel"`
		UsMia struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-mia"`
		ItMil struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"it-mil"`
		CaMtr struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ca-mtr"`
		UsNyc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-nyc"`
		CyNic struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"cy-nic"`
		JpOsa struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"jp-osa"`
		NoOsl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"no-osl"`
		ItPmo struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"it-pmo"`
		FrPar struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"fr-par"`
		AuPer struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"au-per"`
		UsPhx struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-phx"`
		CzPrg struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"cz-prg"`
		MxQro struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"mx-qro"`
		UsRag struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-rag"`
		UsSlc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-slc"`
		UsSjc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-sjc"`
		ClScl struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"cl-scl"`
		BrSao struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"br-sao"`
		UsSea struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-sea"`
		UsUyk struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-uyk"`
		SgSin struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"sg-sin"`
		BgSof struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"bg-sof"`
		NoSvg struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"no-svg"`
		SeSto struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"se-sto"`
		AuSyd struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"au-syd"`
		EeTll struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ee-tll"`
		IlTlv struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"il-tlv"`
		AlTia struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"al-tia"`
		JpTyo struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"jp-tyo"`
		CaTor struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ca-tor"`
		EsVlc struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"es-vlc"`
		CaVan struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ca-van"`
		AtVie struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"at-vie"`
		PlWaw struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"pl-waw"`
		UsWas struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"us-was"`
		HrZag struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"hr-zag"`
		ChZrh struct {
			Country   string  `json:"country"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"ch-zrh"`
	} `json:"locations"`
	Openvpn struct {
		Relays []interface{} `json:"relays"`
		Ports  []struct {
			Port     int    `json:"port"`
			Protocol string `json:"protocol"`
		} `json:"ports"`
	} `json:"openvpn"`
	Wireguard struct {
		Relays []struct {
			Hostname               string   `json:"hostname"`
			Location               string   `json:"location"`
			Active                 bool     `json:"active"`
			Owned                  bool     `json:"owned"`
			Provider               string   `json:"provider"`
			Stboot                 bool     `json:"stboot"`
			Ipv4AddrIn             string   `json:"ipv4_addr_in"`
			IncludeInCountry       bool     `json:"include_in_country"`
			Weight                 int      `json:"weight"`
			PublicKey              string   `json:"public_key"`
			Ipv6AddrIn             string   `json:"ipv6_addr_in"`
			ShadowsocksExtraAddrIn []string `json:"shadowsocks_extra_addr_in,omitempty"`
			Features               struct {
				Quic struct {
					AddrIn []string `json:"addr_in"`
					Token  string   `json:"token"`
					Domain string   `json:"domain"`
				} `json:"quic"`
			} `json:"features,omitempty"`
			Daita bool `json:"daita,omitempty"`
		} `json:"relays"`
		PortRanges            [][]int `json:"port_ranges"`
		ShadowsocksPortRanges [][]int `json:"shadowsocks_port_ranges"`
		Ipv4Gateway           string  `json:"ipv4_gateway"`
		Ipv6Gateway           string  `json:"ipv6_gateway"`
	} `json:"wireguard"`
	Bridge struct {
		Relays []struct {
			Hostname         string `json:"hostname"`
			Location         string `json:"location"`
			Active           bool   `json:"active"`
			Owned            bool   `json:"owned"`
			Provider         string `json:"provider"`
			Stboot           bool   `json:"stboot"`
			Ipv4AddrIn       string `json:"ipv4_addr_in"`
			IncludeInCountry bool   `json:"include_in_country"`
			Weight           int    `json:"weight"`
		} `json:"relays"`
		Shadowsocks []struct {
			Protocol string `json:"protocol"`
			Port     int    `json:"port"`
			Cipher   string `json:"cipher"`
			Password string `json:"password"`
		} `json:"shadowsocks"`
	} `json:"bridge"`
}

type SimpleRelay struct {
	Hostname string
	IPv4Addr string
	// MultihopPort int
	PublicKey string
	CityCode  string
}

type FilteredRelays struct {
	relays []SimpleRelay
	ports  [][]int
}

var currentRelay RelayList

func RelayFilter(client Client, relays RelayList) FilteredRelays {
	locationList := client.Locations
	var WantedRelays FilteredRelays
	WantedRelays.ports = append(WantedRelays.ports, relays.Wireguard.PortRanges...)
	for i := 0; i < len(locationList); i++ {

		json.Unmarshal([]byte(relays.Wireguard.Relays[i].Provider), &currentRelay)
		for x := 0; x < len(relays.Wireguard.Relays); x++ {
			// if country item in locations
			// 		do thing
			if slices.Contains(locationList, relays.Wireguard.Relays[x].Location) {
				var singleRelay SimpleRelay
				singleRelay = SimpleRelay{
					Hostname:  relays.Wireguard.Relays[x].Hostname,
					IPv4Addr:  relays.Wireguard.Relays[x].Ipv4AddrIn,
					PublicKey: relays.Wireguard.Relays[x].PublicKey,
					CityCode:  relays.Wireguard.Relays[x].Location,
				}
				WantedRelays.relays = append(WantedRelays.relays, singleRelay)
			} else {
				continue
			}
		}
	}
		return WantedRelays
}

// will do json parsing and relay filtering
