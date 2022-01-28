package hetzner

// RDNS information about rDNS
type RDNS struct {
	IP  string `json:"ip"`
	Ptr string `json:"ptr"`
}

type dataRDNS struct {
	RDNS *RDNS `json:"rdns"`
}

// RDNSUpdateRequest Requested information for RDNSService.Create and RDNSService.Update
type RDNSUpdateRequest struct {
	IP  string `url:"ip"`
	Ptr string `url:"ptr"`
}
