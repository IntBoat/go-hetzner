package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#subnet

// SubnetService represents a service to work with Subnet.
type SubnetService interface {
	// List Query list of all subnets
	// See: https://robot.your-server.de/doc/webservice/en.html#get-subnet
	List() ([]*Subnet, *http.Response, error)
	// Get Query data of a specific subnet
	// See: https://robot.your-server.de/doc/webservice/en.html#get-subnet-net-ip
	Get(ip string) (*Subnet, *http.Response, error)
	// UpdateWarring Update traffic warning options for a subnet
	// See: https://robot.your-server.de/doc/webservice/en.html#post-subnet-net-ip
	UpdateWarring(req *SubnetUpdateRequest) (*Subnet, *http.Response, error)

	// GetMac Query if it is possible to set a separate MAC address.
	// See: https://robot.your-server.de/doc/webservice/en.html#get-subnet-net-ip-mac
	GetMac(ip string) (*SubnetMac, *http.Response, error)
	// CreateMac Generate a separate MAC address
	// See: https://robot.your-server.de/doc/webservice/en.html#put-subnet-net-ip-mac
	CreateMac(ip string) (*SubnetMac, *http.Response, error)
	// DeleteMac Remove a separate MAC address and set it to the default value
	// (The MAC address of the servers main IP address).
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-subnet-net-ip-mac
	DeleteMac(ip string) (*SubnetMac, *http.Response, error)
}

type SubnetServiceImpl struct {
	client *Client
}

var _ SubnetService = &SubnetServiceImpl{}

func (s *SubnetServiceImpl) List() ([]*Subnet, *http.Response, error) {
	path := "/subnet"

	data := make([]dataSubnet, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Subnet, len(data))
	for i, d := range data {
		a[i] = d.Subnet
	}
	return a, resp, err
}

func (s *SubnetServiceImpl) Get(netIP string) (*Subnet, *http.Response, error) {
	path := fmt.Sprintf("/subnet/%s", netIP)

	data := dataSubnet{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Subnet, resp, err
}

func (s *SubnetServiceImpl) UpdateWarring(req *SubnetUpdateRequest) (*Subnet, *http.Response, error) {
	path := fmt.Sprintf("/subnet/%s", req.NetIP)

	data := dataSubnet{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Subnet, resp, err
}

func (s *SubnetServiceImpl) GetMac(netIP string) (*SubnetMac, *http.Response, error) {
	path := fmt.Sprintf("/subnet/%s/mac", netIP)

	data := dataSubnetMac{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Mac, resp, err
}

func (s *SubnetServiceImpl) CreateMac(netIP string) (*SubnetMac, *http.Response, error) {
	path := fmt.Sprintf("/subnet/%s/mac", netIP)

	data := dataSubnetMac{}
	resp, err := s.client.Call(http.MethodPut, path, nil, &data)
	return data.Mac, resp, err
}

func (s *SubnetServiceImpl) DeleteMac(netIP string) (*SubnetMac, *http.Response, error) {
	path := fmt.Sprintf("/subnet/%s/mac", netIP)

	data := dataSubnetMac{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Mac, resp, err
}
