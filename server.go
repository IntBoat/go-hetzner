package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#server

// ServerService represents a service to work with server service.
type ServerService interface {
	List() ([]*ServerSummary, *http.Response, error)
	Get(serverIP string) (*Server, *http.Response, error)
	Update(req *UpdateServerRequest) (*Server, *http.Response, error)

	GetCancellation(serverIP string) (*Cancellation, *http.Response, error)
	CancelServer(req *CancelServerRequest) (*Cancellation, *http.Response, error)
	WithdrawCancellation(serverIP string) (*http.Response, error)
}

type ServerServiceImpl struct {
	client *Client
}

var _ ServerService = &ServerServiceImpl{}

// List Query data of all servers
// See: https://robot.your-server.de/doc/webservice/en.html#get-server
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

// Get Query server data for a specific server
// See: https://robot.your-server.de/doc/webservice/en.html#get-server-server-number
func (s *ServerServiceImpl) Get(ServerNumber string) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%s", ServerNumber)

	data := dataServer{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Server, resp, err
}

// Update server name for a specific server
// See: https://robot.your-server.de/doc/webservice/en.html#post-server-server-number
func (s *ServerServiceImpl) Update(req *UpdateServerRequest) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%s", req.ServerNumber)

	data := dataServer{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Server, resp, err
}

// GetCancellation Query cancellation data for a server
// See: https://robot.your-server.de/doc/webservice/en.html#get-server-server-number-cancellation
func (s *ServerServiceImpl) GetCancellation(ServerNumber string) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%s/cancellation", ServerNumber)

	data := dataCancellation{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Cancellation, resp, err
}

// CancelServer Cancel a server
// See: https://robot.your-server.de/doc/webservice/en.html#post-server-server-number-cancellation
func (s *ServerServiceImpl) CancelServer(req *CancelServerRequest) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%s/cancellation", req.ServerNumber)

	data := dataCancellation{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cancellation, resp, err
}

// WithdrawCancellation Withdraw a server cancellation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-server-server-number-cancellation
func (s *ServerServiceImpl) WithdrawCancellation(ServerNumber string) (*http.Response, error) {
	path := fmt.Sprintf("/server/%s/cancellation", ServerNumber)

	return s.client.Call(http.MethodDelete, path, nil, nil)
}

// WithdrawOrder Withdraw a server order
// See: https://robot.your-server.de/doc/webservice/en.html#post-server-server-number-reversal
func (s *ServerServiceImpl) WithdrawOrder(req *WithdrawOrderRequest) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%s/reversal", req.ServerNumber)

	data := dataCancellation{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cancellation, resp, err
}
