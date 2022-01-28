package hetzner

import (
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#subnet

// TrafficService represents a service to work with Traffic.
type TrafficService interface {
	// Get Query traffic data for multiple IPs or Subnet
	// See: https://robot.your-server.de/doc/webservice/en.html#post-traffic
	Get(req TrafficRequest) (*Traffic, *http.Response, error)
	// GetByGroup Query traffic data for multiple IPs or Subnet, traffic data is returned not as a sum over the whole
	// interval but grouped by hour, day or month
	// See: https://robot.your-server.de/doc/webservice/en.html#post-traffic
	GetByGroup(req TrafficRequest) (*TrafficGroup, *http.Response, error)
}

type TrafficServiceImpl struct {
	client *Client
}

var _ TrafficService = &TrafficServiceImpl{}

func (s *TrafficServiceImpl) Get(req TrafficRequest) (*Traffic, *http.Response, error) {
	path := "/traffic"

	req.SingleValues = false
	data := dataTraffic{}
	resp, err := s.client.Call(http.MethodGet, path, req, &data)

	return data.Traffic, resp, err
}

func (s *TrafficServiceImpl) GetByGroup(req TrafficRequest) (*TrafficGroup, *http.Response, error) {
	path := "/traffic"

	req.SingleValues = true
	data := dataTrafficGroup{}
	resp, err := s.client.Call(http.MethodGet, path, req, &data)

	return data.Traffic, resp, err
}
