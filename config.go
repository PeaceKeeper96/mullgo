package mullvad

import "fmt"

// WireguardConfig is the details of the generated config file.
type WireGuardConfig struct {
	// Name of the file. E.g. 'mlvd-selse2'
	Name string
	// Content is the content of the generate configuration file.
	Content string
}

func (c *Client) GenerateConfig(Relay []SimpleRelay) (WireGuardConfig, error) {
	if len(Relay) == 0 {
		return WireGuardConfig{}, fmt.Errorf("test")
	}
	return WireGuardConfig{Name: "test", Content: "test123"}, nil
}
// will do wireguard config file generation - seperate to keygen

func GenerateConfig(c Client, Relay []SimpleRelay) (WireGuardConfig, error) {
	return c.GenerateConfig(Relay)
}