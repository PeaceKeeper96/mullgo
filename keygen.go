package mullvad

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)


func Keygen() (Pubkey, privkey string) {
	fmt.Println("\nKeygen")
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return
	}

	return key.String(), key.PublicKey().String()
}
// wireguard key generation