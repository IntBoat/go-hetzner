package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#failover

// FailoverService represents a service to work with failover ips.
type FailoverService interface {
	// List Query failover data for all servers
	// See: https://robot.your-server.de/doc/webservice/en.html#get-failover
	List() ([]*Failover, *http.Response, error)
	// Get Query specific failover IP address data
	// See: https://robot.your-server.de/doc/webservice/en.html#get-failover-failover-ip
	Get(failoverIP string) (*Failover, *http.Response, error)
	// Switch traffic of failoverIP to the server with ip newActiveIP.
	// When successful, returns updated information about the failover IP
	// See: https://robot.your-server.de/doc/webservice/en.html#post-failover-failover-ip
	Switch(req FailoverSwitchRequest) (*Failover, *http.Response, error)
	// Delete the routing of a failover IP
	// When successful, returns updated information about the failover IP
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-failover-failover-ip
	Delete(failoverIP string) (*Failover, *http.Response, error)
}

type FailoverServiceImpl struct {
	client *Client
}

var _ FailoverService = &FailoverServiceImpl{}

func (s *FailoverServiceImpl) List() ([]*Failover, *http.Response, error) {
	path := "/failover"

	data := make([]dataFailover, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Failover, len(data))
	for i, d := range data {
		a[i] = d.Failover
	}
	return a, resp, err
}

func (s *FailoverServiceImpl) Get(failoverIP string) (*Failover, *http.Response, error) {
	path := fmt.Sprintf("/failover/%s", failoverIP)

	data := dataFailover{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Failover, resp, err
}

func (s *FailoverServiceImpl) Switch(req FailoverSwitchRequest) (*Failover, *http.Response, error) {
	path := fmt.Sprintf("/failover/%s", req.FailoverIP)

	data := dataFailover{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Failover, resp, err
}

func (s *FailoverServiceImpl) Delete(failoverIP string) (*Failover, *http.Response, error) {
	path := fmt.Sprintf("/failover/%s", failoverIP)

	data := dataFailover{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Failover, resp, err
}
