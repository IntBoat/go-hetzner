package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#boot-configuration

// BootService represents a service to work with boot service.
type BootService interface {
	GetBoot(serverNumber int) (*Boot, *http.Response, error)

	GetRescue(serverNumber int) (*Rescue, *http.Response, error)
	ActivateRescue(req *ActivateRescueRequest) (*Rescue, *http.Response, error)
	DeactivateRescue(serverNumber int) (*Rescue, *http.Response, error)
	GetRescueLast(serverNumber int) (*Rescue, *http.Response, error)

	GetLinux(serverNumber int) (*Linux, *http.Response, error)
	ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error)
	DeactivateLinux(serverNumber int) (*Linux, *http.Response, error)
	GetLinuxLast(serverNumber int) (*Linux, *http.Response, error)

	GetVnc(serverNumber int) (*Vnc, *http.Response, error)
	ActivateVnc(req *ActivateVncRequest) (*Vnc, *http.Response, error)
	DeactivateVnc(serverNumber int) (*Vnc, *http.Response, error)

	GetWindows(serverNumber int) (*Windows, *http.Response, error)
	ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error)
	DeactivateWindows(serverNumber int) (*Windows, *http.Response, error)

	GetPlesk(serverNumber int) (*Plesk, *http.Response, error)
	ActivatePlesk(req *ActivatePleskRequest) (*Plesk, *http.Response, error)
	DeactivatePlesk(serverNumber int) (*Plesk, *http.Response, error)

	GetCPanel(serverNumber int) (*Cpanel, *http.Response, error)
	ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error)
	DeactivateCPanel(serverNumber int) (*Cpanel, *http.Response, error)
}

type BootServiceImpl struct {
	client *Client
}

var _ BootService = &BootServiceImpl{}

// GetBoot Query the current boot configuration status for a server.
// There can be only one configuration active at any time for one server.
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number
func (s *BootServiceImpl) GetBoot(serverNumber int) (*Boot, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d", serverNumber)

	data := dataBoot{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Boot, resp, err
}

// GetRescue Query boot  for the Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue
func (s *BootServiceImpl) GetRescue(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

// ActivateRescue Activate Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-rescue
func (s *BootServiceImpl) ActivateRescue(req *ActivateRescueRequest) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", req.ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Rescue, resp, err
}

// DeactivateRescue Deactivate Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-rescue
func (s *BootServiceImpl) DeactivateRescue(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Rescue, resp, err
}

// GetRescueLast Show data of last rescue activation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue-last
func (s *BootServiceImpl) GetRescueLast(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue/last", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

// GetLinux Query boot  for the Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux
func (s *BootServiceImpl) GetLinux(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}

// ActivateLinux Activate Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-linux
func (s *BootServiceImpl) ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", req.ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Linux, resp, err
}

// DeactivateLinux Deactivate Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-linux
func (s *BootServiceImpl) DeactivateLinux(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Linux, resp, err
}

// GetLinuxLast Show data of last Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux-last
func (s *BootServiceImpl) GetLinuxLast(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux/last", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}

// GetVnc Query boot  for the VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-vnc
func (s *BootServiceImpl) GetVnc(serverNumber int) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", serverNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Vnc, resp, err
}

// ActivateVnc Activate VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-vnc
func (s *BootServiceImpl) ActivateVnc(req *ActivateVncRequest) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", req.ServerNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Vnc, resp, err
}

// DeactivateVnc Deactivate VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
func (s *BootServiceImpl) DeactivateVnc(serverNumber int) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", serverNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Vnc, resp, err
}

// GetWindows Query boot  for the windows installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
func (s *BootServiceImpl) GetWindows(serverNumber int) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", serverNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Windows, resp, err
}

// ActivateWindows Activate Windows installation.
// You need to order the Windows addon for the server via the Robot web panel first.
// After a reboot, the installation will start, and all data on the server will be deleted.
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-windows
func (s *BootServiceImpl) ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", req.ServerNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Windows, resp, err
}

// DeactivateWindows Deactivate Windows installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
func (s *BootServiceImpl) DeactivateWindows(serverNumber int) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", serverNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Windows, resp, err
}

// GetPlesk Query boot  for the Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
func (s *BootServiceImpl) GetPlesk(serverNumber int) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", serverNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Plesk, resp, err
}

// ActivatePlesk Activate Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-plesk
func (s *BootServiceImpl) ActivatePlesk(req *ActivatePleskRequest) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", req.ServerNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Plesk, resp, err
}

// DeactivatePlesk Deactivate Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-plesk
func (s *BootServiceImpl) DeactivatePlesk(serverNumber int) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", serverNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Plesk, resp, err
}

// GetCPanel Query boot  for the cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-ip-cpanel
func (s *BootServiceImpl) GetCPanel(serverNumber int) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", serverNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Cpanel, resp, err
}

// ActivateCPanel Activate cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-cpanel
func (s *BootServiceImpl) ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", req.ServerNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cpanel, resp, err
}

// DeactivateCPanel Deactivate cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-cpanel
func (s *BootServiceImpl) DeactivateCPanel(serverNumber int) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", serverNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Cpanel, resp, err
}
