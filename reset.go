package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#reset

// ResetService represents a service to work with Reset ordering.
type ResetService interface {
	// List Query reset options for all servers
	// See: https://robot.your-server.de/doc/webservice/en.html#get-reset
	List() ([]*Reset, *http.Response, error)
	// Get Query reset options for a specific server
	// See: https://robot.your-server.de/doc/webservice/en.html#get-reset-server-number
	Get(serverNumber int) (*Reset, *http.Response, error)
	// Reset Execute reset on specific server
	// See: https://robot.your-server.de/doc/webservice/en.html#post-reset-server-number
	Reset(req *ResetCreateRequest) (*Reset, *http.Response, error)
}

type ResetServiceImpl struct {
	client *Client
}

var _ ResetService = &ResetServiceImpl{}

func (s *ResetServiceImpl) List() ([]*Reset, *http.Response, error) {
	path := "/reset"

	data := make([]dataReset, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Reset, len(data))
	for i, d := range data {
		a[i] = d.Reset
	}
	return a, resp, err
}

func (s *ResetServiceImpl) Get(serverNumber int) (*Reset, *http.Response, error) {
	path := fmt.Sprintf("/reset/%d", serverNumber)

	data := dataReset{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Reset, resp, err
}

func (s *ResetServiceImpl) Reset(req *ResetCreateRequest) (*Reset, *http.Response, error) {
	path := fmt.Sprintf("/reset/%d", req.ServerNumber)

	data := dataResetCreateResponse{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)

	out := &Reset{
		ServerIP:     data.Reset.ServerIP,
		ServerNumber: data.Reset.ServerNumber,
		Type:         []string{data.Reset.Type},
	}
	return out, resp, err
}
