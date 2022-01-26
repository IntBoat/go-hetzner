package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#wake-on-lan

// WOLService represents a service to work with Wake on LAN service.
type WOLService interface {
	Create(serverIP string) (*WOL, *http.Response, error)
	Get(serverIP string) (*WOL, *http.Response, error)
}

type WOLServiceImpl struct {
	client *Client
}

var _ WOLService = &WOLServiceImpl{}

// Get Query Wake On LAN data
// See: https://robot.your-server.de/doc/webservice/en.html#get-wol-server-number
func (s *WOLServiceImpl) Get(serverNumber string) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%s", serverNumber)

	data := dataWol{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.WOL, resp, err
}

// Create Send Wake On LAN packet to server
// See: https://robot.your-server.de/doc/webservice/en.html#post-wol-server-number
func (s *WOLServiceImpl) Create(serverNumber string) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%s", serverNumber)

	data := dataWol{}
	resp, err := s.client.Call(http.MethodPost, path, nil, &data)
	return data.WOL, resp, err
}
