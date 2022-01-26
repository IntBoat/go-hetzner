package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#ssh-keys

// SSHKeyService represents a service to work with SSH service.
type SSHKeyService interface {
	List() ([]*SSHKey, *http.Response, error)
	Create(req *SSHKeyCreateRequest) (*SSHKey, *http.Response, error)
	Get(fingerprint string) (*SSHKey, *http.Response, error)
	Update(req *SSHKeyUpdateRequest) (*SSHKey, *http.Response, error)
	Delete(fingerprint string) (*http.Response, error)
}

type SSHKeyServiceImpl struct {
	client *Client
}

var _ SSHKeyService = &SSHKeyServiceImpl{}

// List Query all SSH keys
// See: https://robot.your-server.de/doc/webservice/en.html#get-key
func (s *SSHKeyServiceImpl) List() ([]*SSHKey, *http.Response, error) {
	path := "/key"

	data := make([]dataKey, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*SSHKey, len(data))
	for i, d := range data {
		a[i] = d.Key
	}
	return a, resp, err
}

// Create Add a new SSH key. Once the key is successfully added, the status code 201 CREATED is returned.
// See: https://robot.your-server.de/doc/webservice/en.html#post-key
func (s *SSHKeyServiceImpl) Create(req *SSHKeyCreateRequest) (*SSHKey, *http.Response, error) {
	path := "/key"

	data := dataKey{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Key, resp, err
}

// Get Query a specific SSH key
// See: https://robot.your-server.de/doc/webservice/en.html#get-key-fingerprint
func (s *SSHKeyServiceImpl) Get(fingerprint string) (*SSHKey, *http.Response, error) {
	path := fmt.Sprintf("/key/%v", fingerprint)

	data := dataKey{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Key, resp, err
}

// Update the key name
// See: https://robot.your-server.de/doc/webservice/en.html#post-key-fingerprint
func (s *SSHKeyServiceImpl) Update(req *SSHKeyUpdateRequest) (*SSHKey, *http.Response, error) {
	path := fmt.Sprintf("/key/%v", req.Fingerprint)

	data := dataKey{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Key, resp, err
}

// Delete Remove public key
// See: https://robot.your-server.de/doc/webservice/en.html#delete-key-fingerprint
func (s *SSHKeyServiceImpl) Delete(fingerprint string) (*http.Response, error) {
	path := fmt.Sprintf("/key/%v", fingerprint)
	return s.client.Call(http.MethodDelete, path, nil, nil)
}
