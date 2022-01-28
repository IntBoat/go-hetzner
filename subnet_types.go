package hetzner

// Subnet contains information about a Subnet
type Subnet struct {
	IP              string `json:"ip"`
	Mask            int    `json:"mask"`
	Gateway         string `json:"gateway"`
	ServerIp        string `json:"server_ip"`
	ServerNumber    int    `json:"server_number"`
	Failover        bool   `json:"failover"`
	Locked          bool   `json:"locked"`
	TrafficWarnings bool   `json:"traffic_warnings"`
	TrafficHourly   int    `json:"traffic_hourly"`
	TrafficDaily    int    `json:"traffic_daily"`
	TrafficMonthly  int    `json:"traffic_monthly"`
}

// SubnetMac contains information about a Subnet Mac
type SubnetMac struct {
	IP          string              `json:"ip"`
	Mask        int                 `json:"mask"`
	Mac         string              `json:"mac"`
	PossibleMac []map[string]string `json:"possible_mac"`
}

// SubnetUpdateRequest Requested information for SubnetService.UpdateWarring
type SubnetUpdateRequest struct {
	NetIP           string // Net IP
	TrafficWarnings string `url:"traffic_warnings"` // Enable/disable traffic warnings (true, false)
	TrafficHourly   string `url:"traffic_hourly"`   // Hourly traffic limit in MB
	TrafficDaily    string `url:"traffic_daily"`    // Daily traffic limit in MB
	TrafficMonthly  string `url:"traffic_monthly"`  // Monthly traffic limit in GB
}

type dataSubnet struct {
	Subnet *Subnet `json:"subnet"`
}

type dataSubnetMac struct {
	Mac *SubnetMac `json:"mac"`
}
