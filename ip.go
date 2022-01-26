package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#ip

// IpService represents a service to work with IP.
type IpService interface {
	List() ([]*IP, *http.Response, error)
	Get(ip string) (*IpSummary, *http.Response, error)
	UpdateWarring(req *IpUpdateRequest) (*IpSummary, *http.Response, error)

	GetMac(ip string) (*Mac, *http.Response, error)
	CreateMac(ip string) (*Mac, *http.Response, error)
	DeleteMac(ip string) (*Mac, *http.Response, error)
}

type IpServiceImpl struct {
	client *Client
}

var _ IpService = &IpServiceImpl{}

// List Query list of all single IP addresses
// See: https://robot.your-server.de/doc/webservice/en.html#get-ip
func (s *IpServiceImpl) List() ([]*IP, *http.Response, error) {
	path := "/ip"

	data := make([]dataIp, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*IP, len(data))
	for i, d := range data {
		a[i] = d.IP
	}
	return a, resp, err
}

// Get Query data for a specific IP address
// See: https://robot.your-server.de/doc/webservice/en.html#get-ip-ip
func (s *IpServiceImpl) Get(ip string) (*IpSummary, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", ip)

	data := dataIpSummary{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.IpSummary, resp, err
}

// UpdateWarring update traffic warning options for an IP address
// See: https://robot.your-server.de/doc/webservice/en.html#post-ip-ip
func (s *IpServiceImpl) UpdateWarring(req *IpUpdateRequest) (*IpSummary, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", req.IP)

	data := dataIpSummary{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.IpSummary, resp, err
}

// GetMac Query if it is possible to set a separate MAC address. Returns the MAC address if it is set.
// See: https://robot.your-server.de/doc/webservice/en.html#get-ip-ip-mac
func (s *IpServiceImpl) GetMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Mac, resp, err
}

// CreateMac Generate a separate MAC address
// See: https://robot.your-server.de/doc/webservice/en.html#put-ip-ip-mac
func (s *IpServiceImpl) CreateMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodPut, path, nil, &data)
	return data.Mac, resp, err
}

// DeleteMac Remove a separate MAC address
// See: https://robot.your-server.de/doc/webservice/en.html#delete-ip-ip-mac
func (s *IpServiceImpl) DeleteMac(ip string) (*Mac, *http.Response, error) {
	path := fmt.Sprintf("/ip/%s/mac", ip)

	data := dataMac{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Mac, resp, err
}
