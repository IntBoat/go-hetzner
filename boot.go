package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#boot-configuration

// BootService represents a service to work with boot service.
type BootService interface {
	// GetBoot Query the current boot configuration status for a server.
	// There can be only one configuration active at any time for one server.
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number
	GetBoot(serverNumber int) (*Boot, *http.Response, error)

	// GetRescue Query boot  for the Rescue System
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue
	GetRescue(serverNumber int) (*Rescue, *http.Response, error)
	// ActivateRescue Activate Rescue System
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-rescue
	ActivateRescue(req *ActivateRescueRequest) (*Rescue, *http.Response, error)
	// DeactivateRescue Deactivate Rescue System
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-rescue
	DeactivateRescue(serverNumber int) (*Rescue, *http.Response, error)
	// GetRescueLast Show data of last rescue activation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue-last
	GetRescueLast(serverNumber int) (*Rescue, *http.Response, error)

	// GetLinux Query boot  for the Linux installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux
	GetLinux(serverNumber int) (*Linux, *http.Response, error)
	// ActivateLinux Activate Linux installation
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-linux
	ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error)
	// DeactivateLinux Deactivate Linux installation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-linux
	DeactivateLinux(serverNumber int) (*Linux, *http.Response, error)
	// GetLinuxLast Show data of last Linux installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux-last
	GetLinuxLast(serverNumber int) (*Linux, *http.Response, error)

	// GetVnc Query boot  for the VNC installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-vnc
	GetVnc(serverNumber int) (*Vnc, *http.Response, error)
	// ActivateVnc Activate VNC installation
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-vnc
	ActivateVnc(req *ActivateVncRequest) (*Vnc, *http.Response, error)
	// DeactivateVnc Deactivate VNC installation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
	DeactivateVnc(serverNumber int) (*Vnc, *http.Response, error)

	// GetWindows Query boot  for the windows installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
	GetWindows(serverNumber int) (*Windows, *http.Response, error)
	// ActivateWindows Activate Windows installation.
	// You need to order the Windows addon for the server via the Robot web panel first.
	// After a reboot, the installation will start, and all data on the server will be deleted.
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-windows
	ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error)
	// DeactivateWindows Deactivate Windows installation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
	DeactivateWindows(serverNumber int) (*Windows, *http.Response, error)

	// GetPlesk Query boot  for the Plesk installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
	GetPlesk(serverNumber int) (*Plesk, *http.Response, error)
	// ActivatePlesk Activate Plesk installation
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-plesk
	ActivatePlesk(req *ActivatePleskRequest) (*Plesk, *http.Response, error)
	// DeactivatePlesk Deactivate Plesk installation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-plesk
	DeactivatePlesk(serverNumber int) (*Plesk, *http.Response, error)

	// GetCPanel Query boot  for the cPanel installation
	// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-ip-cpanel
	GetCPanel(serverNumber int) (*Cpanel, *http.Response, error)
	// ActivateCPanel Activate cPanel installation
	// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-cpanel
	ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error)
	// DeactivateCPanel Deactivate cPanel installation
	// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-cpanel
	DeactivateCPanel(serverNumber int) (*Cpanel, *http.Response, error)
}

type BootServiceImpl struct {
	client *Client
}

var _ BootService = &BootServiceImpl{}

func (s *BootServiceImpl) GetBoot(serverNumber int) (*Boot, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d", serverNumber)

	data := dataBoot{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Boot, resp, err
}

func (s *BootServiceImpl) GetRescue(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

func (s *BootServiceImpl) ActivateRescue(req *ActivateRescueRequest) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", req.ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Rescue, resp, err
}

func (s *BootServiceImpl) DeactivateRescue(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Rescue, resp, err
}

func (s *BootServiceImpl) GetRescueLast(serverNumber int) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/rescue/last", serverNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

func (s *BootServiceImpl) GetLinux(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}
func (s *BootServiceImpl) ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", req.ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Linux, resp, err
}

func (s *BootServiceImpl) DeactivateLinux(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Linux, resp, err
}
func (s *BootServiceImpl) GetLinuxLast(serverNumber int) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/linux/last", serverNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}

func (s *BootServiceImpl) GetVnc(serverNumber int) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", serverNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Vnc, resp, err
}

func (s *BootServiceImpl) ActivateVnc(req *ActivateVncRequest) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", req.ServerNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Vnc, resp, err
}

func (s *BootServiceImpl) DeactivateVnc(serverNumber int) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/vnc", serverNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Vnc, resp, err
}

func (s *BootServiceImpl) GetWindows(serverNumber int) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", serverNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Windows, resp, err
}

func (s *BootServiceImpl) ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", req.ServerNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Windows, resp, err
}

func (s *BootServiceImpl) DeactivateWindows(serverNumber int) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/windows", serverNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Windows, resp, err
}

func (s *BootServiceImpl) GetPlesk(serverNumber int) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", serverNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Plesk, resp, err
}

func (s *BootServiceImpl) ActivatePlesk(req *ActivatePleskRequest) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", req.ServerNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Plesk, resp, err
}

func (s *BootServiceImpl) DeactivatePlesk(serverNumber int) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/plesk", serverNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Plesk, resp, err
}

func (s *BootServiceImpl) GetCPanel(serverNumber int) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", serverNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Cpanel, resp, err
}
func (s *BootServiceImpl) ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", req.ServerNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cpanel, resp, err
}

func (s *BootServiceImpl) DeactivateCPanel(serverNumber int) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%d/cpanel", serverNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Cpanel, resp, err
}
