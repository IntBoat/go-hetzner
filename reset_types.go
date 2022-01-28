package hetzner

// Reset contains information about reset service
type Reset struct {
	ServerIP        string   `json:"server_ip"`
	ServerIpv6Net   string   `json:"server_ipv6_net"`
	ServerNumber    int      `json:"server_number"`
	Type            []string `json:"type"`
	OperatingStatus string   `json:"operating_status"`
}

// ResetRequest requested information for ResetService.Reset
type ResetRequest struct {
	ServerNumber int
	// Reset type
	// Send CTRL+ALT+DEL to the server: sw
	// Press hardware reset button: hw
	// order a manual hardware reset: man
	Type string `url:"type"`
}

// ResetResponse contains information about current reset status
type ResetResponse struct {
	ServerIP      string `json:"server_ip"`
	ServerIpv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
	Type          string `json:"type"`
}

type dataReset struct {
	Reset *Reset `json:"reset"`
}

type dataResetResponse struct {
	Reset *ResetResponse `json:"reset"`
}
