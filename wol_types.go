package hetzner

// WOL contains information about WOL
type WOL struct {
	ServerIp      string `json:"server_ip"`
	ServerIpv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
}

type dataWol struct {
	WOL *WOL `json:"wol"`
}
