package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#wake-on-lan

// WOLService represents a service to work with Wake on LAN service.
type WOLService interface {
	// Get Query Wake On LAN data
	// See: https://robot.your-server.de/doc/webservice/en.html#get-wol-server-number
	Get(serverNumber int) (*WOL, *http.Response, error)
	// Create Send Wake On LAN packet to server
	// See: https://robot.your-server.de/doc/webservice/en.html#post-wol-server-number
	Create(serverNumber int) (*WOL, *http.Response, error)
}

type WOLServiceImpl struct {
	client *Client
}

var _ WOLService = &WOLServiceImpl{}

func (s *WOLServiceImpl) Get(serverNumber int) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%d", serverNumber)

	data := dataWol{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.WOL, resp, err
}

func (s *WOLServiceImpl) Create(serverNumber int) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%d", serverNumber)

	data := dataWol{}
	resp, err := s.client.Call(http.MethodPost, path, nil, &data)
	return data.WOL, resp, err
}
