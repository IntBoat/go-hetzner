package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#server

// ServerService represents a service to work with server service.
type ServerService interface {
	// List Query data of all servers
	// See: https://robot.your-server.de/doc/webservice/en.html#get-server
	List() ([]*ServerSummary, *http.Response, error)
	// Get Query server data for a specific server
	// See: https://robot.your-server.de/doc/webservice/en.html#get-server-server-number
	Get(serverNumber int) (*Server, *http.Response, error)
	// Update server name for a specific server
	// See: https://robot.your-server.de/doc/webservice/en.html#post-server-server-number
	Update(req *UpdateServerRequest) (*Server, *http.Response, error)

	// GetCancellation Query cancellation data for a server
	// See: https://robot.your-server.de/doc/webservice/en.html#get-server-server-number-cancellation
	GetCancellation(serverNumber int) (*Cancellation, *http.Response, error)
	// CancelServer Cancel a server
	// See: https://robot.your-server.de/doc/webservice/en.html#post-server-server-number-cancellation
	CancelServer(req *CancelServerRequest) (*CancellationSingle, *http.Response, error)
	// WithdrawCancellation Withdraw a server cancellation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-server-server-number-cancellation
	WithdrawCancellation(serverNumber int) (*http.Response, error)
}

type ServerServiceImpl struct {
	client *Client
}

var _ ServerService = &ServerServiceImpl{}

func (s *ServerServiceImpl) List() ([]*ServerSummary, *http.Response, error) {
	path := "/server"

	data := make([]dataServerSummary, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*ServerSummary, len(data))
	for i, d := range data {
		a[i] = d.Server
	}
	return a, resp, err
}

func (s *ServerServiceImpl) Get(serverNumber int) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%d", serverNumber)

	data := dataServer{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Server, resp, err
}

func (s *ServerServiceImpl) Update(req *UpdateServerRequest) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%d", req.ServerNumber)

	data := dataServer{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Server, resp, err
}

func (s *ServerServiceImpl) GetCancellation(serverNumber int) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%d/cancellation", serverNumber)

	data := dataCancellation{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Cancellation, resp, err
}

func (s *ServerServiceImpl) CancelServer(req *CancelServerRequest) (*CancellationSingle, *http.Response, error) {
	path := fmt.Sprintf("/server/%d/cancellation", req.ServerNumber)

	data := dataCancellationSingle{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cancellation, resp, err
}

func (s *ServerServiceImpl) WithdrawCancellation(serverNumber int) (*http.Response, error) {
	path := fmt.Sprintf("/server/%d/cancellation", serverNumber)

	return s.client.Call(http.MethodDelete, path, nil, nil)
}

func (s *ServerServiceImpl) WithdrawOrder(req *WithdrawOrderRequest) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%s/reversal", req.ServerNumber)

	data := dataCancellation{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cancellation, resp, err
}
