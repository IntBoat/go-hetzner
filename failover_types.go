package hetzner

// Failover contains information about failover ip
type Failover struct {
	IP             string `json:"ip"`
	Netmask        string `json:"netmask"`
	ServerIP       string `json:"server_ip"`
	ServerNumber   int    `json:"server_number"`
	ActiveServerIP string `json:"active_server_ip"`
}

// FailoverSwitchRequest describes parameters of POST /failover API call
type FailoverSwitchRequest struct {
	FailoverIP     string
	ActiveServerIP string `url:"active_server_ip"` // Main IP address of the server where the failover IP should be routed to.
}

type dataFailover struct {
	Failover *Failover `json:"failover"`
}
