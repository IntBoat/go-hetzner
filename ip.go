package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#ip

// IPService represents a service to work with IP.
type IPService interface {
	// List Query list of all single IP addresses
	// See: https://robot.your-server.de/doc/webservice/en.html#get-ip
	List() ([]*IP, *http.Response, error)
	// Get Query data for a specific IP address
	// See: https://robot.your-server.de/doc/webservice/en.html#get-ip-ip
	Get(ip string) (*IPSummary, *http.Response, error)
	// UpdateWarring update traffic warning options for an IP address
	// See: https://robot.your-server.de/doc/webservice/en.html#post-ip-ip
	UpdateWarring(req *IPUpdateRequest) (*IPSummary, *http.Response, error)
	// GetMac Query if it is possible to set a separate MAC address. Returns the MAC address if it is set.
	// See: https://robot.your-server.de/doc/webservice/en.html#get-ip-ip-mac
	GetMac(ip string) (*Mac, *http.Response, error)
	// CreateMac Generate a separate MAC address
	// See: https://robot.your-server.de/doc/webservice/en.html#put-ip-ip-mac
	CreateMac(ip string) (*Mac, *http.Response, error)
	// DeleteMac Remove a separate MAC address
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-ip-ip-mac
	DeleteMac(ip string) (*Mac, *http.Response, error)
}

type IPServiceImpl struct {
	client *Client
}

var _ IPService = &IPServiceImpl{}

func (s *IPServiceImpl) List() ([]*IP, *http.Response, error) {
	path := "/ip"

	data := make([]dataIp, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*IP, len(data))
	for i, d := range data {
		a[i] = d.IP
	}
	return a, resp, err
}

func (s *IPServiceImpl) Get(ip string) (*IPSummary, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", ip)

	data := dataIpSummary{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.IPSummary, resp, err
}

func (s *IPServiceImpl) UpdateWarring(req *IPUpdateRequest) (*IPSummary, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", req.IP)

	data := dataIpSummary{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.IPSummary, resp, err
}
func (s *IPServiceImpl) GetMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s/mac", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Mac, resp, err
}

func (s *IPServiceImpl) CreateMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s/mac", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodPut, path, nil, &data)
	return data.Mac, resp, err
}

func (s *IPServiceImpl) DeleteMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s/mac", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Mac, resp, err
}
