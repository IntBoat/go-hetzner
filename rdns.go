package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#delete-rdns-ip

// RDNSService represents a service to work with Reverse DNS.
type RDNSService interface {
	List() ([]*RDNS, *http.Response, error)
	Get(ip string) (*RDNS, *http.Response, error)
	Create(req *RDNSUpdateRequest) (*RDNS, *http.Response, error)
	Update(req *RDNSUpdateRequest) (*RDNS, *http.Response, error)
	Delete(ip string) (*RDNS, *http.Response, error)
}

type RDNSServiceImpl struct {
	client *Client
}

var _ RDNSService = &RDNSServiceImpl{}

// List Query all rDNS entries
// See: https://robot.your-server.de/doc/webservice/en.html#get-rdns
func (s *RDNSServiceImpl) List() ([]*RDNS, *http.Response, error) {
	path := "/rdns"

	data := make([]dataRDNS, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*RDNS, len(data))
	for i, d := range data {
		a[i] = d.RDNS
	}
	return a, resp, err
}

// Get Query the current reverse DNS entry for one IP address
// See: https://robot.your-server.de/doc/webservice/en.html#get-ip-ip
func (s *RDNSServiceImpl) Get(ip string) (*RDNS, *http.Response, error) {
	path := fmt.Sprintf("/rdns/%s", ip)

	data := dataRDNS{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.RDNS, resp, err
}

// Create new reverse DNS entry for one IP address.
// Once the reverse DNS entry is successfully created, the status code 201 CREATED is returned.
// See: https://robot.your-server.de/doc/webservice/en.html#put-rdns-ip
func (s *RDNSServiceImpl) Create(req *RDNSUpdateRequest) (*RDNS, *http.Response, error) {
	path := fmt.Sprintf("/rdns/%s", req.IP)

	data := dataRDNS{}
	resp, err := s.client.Call(http.MethodPut, path, req, &data)
	return data.RDNS, resp, err
}

// Update create or Update a reverse DNS entry for one IP.
// Once the reverse DNS entry is successfully created, the status code is set to 201 created.
// On successful updates, the status code is 200 OK.
// See: https://robot.your-server.de/doc/webservice/en.html#post-rdns-ip
func (s *RDNSServiceImpl) Update(req *RDNSUpdateRequest) (*RDNS, *http.Response, error) {
	path := fmt.Sprintf("/rdns/%s", req.IP)

	data := dataRDNS{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.RDNS, resp, err
}

// Delete reverse DNS entry for one IP
// See: https://robot.your-server.de/doc/webservice/en.html#delete-rdns-ip
func (s *RDNSServiceImpl) Delete(ip string) (*RDNS, *http.Response, error) {
	path := fmt.Sprintf("/rdns/%s", ip)

	data := dataRDNS{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.RDNS, resp, err
}
