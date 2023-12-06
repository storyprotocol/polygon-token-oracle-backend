package config

type Config struct {
	Debug   *DebugConfig
	Event   *EventConfig
	Polygon *PolygonConfig
	Wallet  *WalletConfig
}

type DebugConfig struct {
	Enable  bool
	Verbose bool
}

type EventConfig struct {
	Endpoint             string
	PolygonTokenHookAddr string
}

type PolygonConfig struct {
	Endpoint string
}

type WalletConfig struct {
	PrivateKey string
}
