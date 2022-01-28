package hetzner

// IP contains information about an IP
type IP struct {
	IP              string `json:"ip"`
	Locked          bool   `json:"locked"`
	SeparateMac     string `json:"separate_mac"`
	ServerIp        string `json:"server_ip"`
	ServerNumber    int    `json:"server_number"`
	TrafficDaily    int    `json:"traffic_daily"`
	TrafficHourly   int    `json:"traffic_hourly"`
	TrafficMonthly  int    `json:"traffic_monthly"`
	TrafficWarnings bool   `json:"traffic_warnings"`
}

// IpSummary contains information about an IP with broadcast, gateway and mask information
type IpSummary struct {
	IP
	Broadcast string `json:"broadcast"`
	Gateway   string `json:"gateway"`
	Mask      int    `json:"mask"`
}

// IpUpdateRequest Requested information for IpService.UpdateWarring
type IpUpdateRequest struct {
	IP              string
	TrafficWarnings string `url:"traffic_warnings"` // Enable/disable traffic warnings (true, false)
	TrafficHourly   string `url:"traffic_hourly"`   // Hourly traffic limit in MB
	TrafficDaily    string `url:"traffic_daily"`    // Daily traffic limit in MB
	TrafficMonthly  string `url:"traffic_monthly"`  // Monthly traffic limit in GB
}

// Mac contains information about a Mac address
type Mac struct {
	IP  string `json:"ip"`
	Mac string `json:"mac"`
}

type dataIp struct {
	IP *IP `json:"ip"`
}

type dataIpSummary struct {
	IpSummary *IpSummary `json:"ip"`
}

type dataMac struct {
	Mac *Mac `json:"mac"`
}
