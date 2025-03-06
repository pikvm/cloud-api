package api_models

type AvailableProxy struct {
	Endpoint string `json:"endpoint"`
}

type AvailableProxies struct {
	Proxies []AvailableProxy `json:"available_proxies"`
}
