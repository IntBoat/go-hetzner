package hetzner

type Reset struct {
	ServerIP        string   `json:"server_ip"`
	ServerIpv6Net   string   `json:"server_ipv6_net"`
	ServerNumber    int      `json:"server_number"`
	Type            []string `json:"type"`
	OperatingStatus string   `json:"operating_status"`
}

type ResetCreateRequest struct {
	ServerNumber string
	Type         string `url:"type"`
}

type ResetCreateResponse struct {
	ServerIP      string `json:"server_ip"`
	ServerIpv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
	Type          string `json:"type"`
}

type dataReset struct {
	Reset *Reset `json:"reset"`
}

type dataResetCreateResponse struct {
	Reset *ResetCreateResponse `json:"reset"`
}
