package hetzner

type RDNS struct {
	IP  string `json:"ip"`
	Ptr string `json:"ptr"`
}

type dataRDNS struct {
	RDNS *RDNS `json:"rdns"`
}

type RDNSUpdateRequest struct {
	IP  string `url:"ip"`
	Ptr string `url:"ptr"`
}
