package mullgo

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// wireguard key generation
func Keygen() (pubkey string, privkey string, error error) {
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", "", fmt.Errorf("Keygen: failed to mgenerate privatekey: %w", err)
	}

	return key.PublicKey().String(), key.String(), nil
}