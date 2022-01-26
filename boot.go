package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#boot-configuration

// BootService represents a service to work with boot service.
type BootService interface {
	GetBoot(ServerNumber string) (*Boot, *http.Response, error)

	GetRescue(ServerNumber string) (*Rescue, *http.Response, error)
	ActivateRescue(req ActivateRescueRequest) (*Rescue, *http.Response, error)
	DeactivateRescue(ServerNumber string) (*Rescue, *http.Response, error)

	GetRescueLast(ServerNumber string) (*Rescue, *http.Response, error)
	GetLinux(ServerNumber string) (*Linux, *http.Response, error)
	ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error)
	DeactivateLinux(ServerNumber string) (*Linux, *http.Response, error)
	GetLinuxLast(ServerNumber string) (*Linux, *http.Response, error)

	GetVnc(ServerNumber string) (*Vnc, *http.Response, error)
	ActivateVnc(req *ActivateLinuxRequest) (*Vnc, *http.Response, error)
	DeactivateVnc(ServerNumber string) (*Vnc, *http.Response, error)

	GetWindows(ServerNumber string) (*Windows, *http.Response, error)
	ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error)
	DeactivateWindows(ServerNumber string) (*Windows, *http.Response, error)

	GetPlesk(ServerNumber string) (*Plesk, *http.Response, error)
	ActivatePlesk(req *ActivateWindowsRequest) (*Plesk, *http.Response, error)
	DeactivatePlesk(ServerNumber string) (*Plesk, *http.Response, error)

	GetCPanel(ServerNumber string) (*Cpanel, *http.Response, error)
	ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error)
	DeactivateCPanel(ServerNumber string) (*Cpanel, *http.Response, error)
}

type BootServiceImpl struct {
	client *Client
}

var _ BootService = &BootServiceImpl{}

// GetBoot Query the current boot configuration status for a server.
// There can be only one configuration active at any time for one server.
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number
func (s *BootServiceImpl) GetBoot(ServerNumber string) (*Boot, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s", ServerNumber)

	data := dataBoot{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Boot, resp, err
}

// GetRescue Query boot options for the Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue
func (s *BootServiceImpl) GetRescue(ServerNumber string) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s", ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

// ActivateRescue Activate Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-rescue
func (s *BootServiceImpl) ActivateRescue(req ActivateRescueRequest) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/rescue", req.ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Rescue, resp, err
}

// DeactivateRescue Deactivate Rescue System
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-rescue
func (s *BootServiceImpl) DeactivateRescue(ServerNumber string) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/rescue", ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, &data)
	return data.Rescue, resp, err
}

// GetRescueLast Show data of last rescue activation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-rescue-last
func (s *BootServiceImpl) GetRescueLast(ServerNumber string) (*Rescue, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/rescue/last", ServerNumber)

	data := dataRescue{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Rescue, resp, err
}

// GetLinux Query boot options for the Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux
func (s *BootServiceImpl) GetLinux(ServerNumber string) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/linux", ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}

// ActivateLinux Activate Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-linux
func (s *BootServiceImpl) ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/linux", req.ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Linux, resp, err
}

// DeactivateLinux Deactivate Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-linux
func (s *BootServiceImpl) DeactivateLinux(ServerNumber string) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/linux", ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, nil)
	return data.Linux, resp, err
}

// GetLinuxLast Show data of last Linux installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-linux-last
func (s *BootServiceImpl) GetLinuxLast(ServerNumber string) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/linux", ServerNumber)

	data := dataLinux{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Linux, resp, err
}

// GetVnc Query boot options for the VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-vnc
func (s *BootServiceImpl) GetVnc(ServerNumber string) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/vnc", ServerNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Vnc, resp, err
}

// ActivateVnc Activate VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-vnc
func (s *BootServiceImpl) ActivateVnc(req *ActivateLinuxRequest) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/vnc", req.ServerNumber)

	type Data struct {
		Vnc *Vnc `json:"vnc"`
	}
	data := dataVnc{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Vnc, resp, err
}

// DeactivateVnc Deactivate VNC installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
func (s *BootServiceImpl) DeactivateVnc(ServerNumber string) (*Vnc, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/vnc", ServerNumber)

	data := dataVnc{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, nil)
	return data.Vnc, resp, err
}

// GetWindows Query boot options for the windows installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
func (s *BootServiceImpl) GetWindows(ServerNumber string) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/windows", ServerNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Windows, resp, err
}

// ActivateWindows Activate Windows installation.
// You need to order the Windows addon for the server via the Robot web panel first.
// After a reboot, the installation will start, and all data on the server will be deleted.
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-windows
func (s *BootServiceImpl) ActivateWindows(req *ActivateWindowsRequest) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/windows", req.ServerNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Windows, resp, err
}

// DeactivateWindows Deactivate Windows installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-vnc
func (s *BootServiceImpl) DeactivateWindows(ServerNumber string) (*Windows, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/windows", ServerNumber)

	data := dataWindows{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, nil)
	return data.Windows, resp, err
}

// GetPlesk Query boot options for the Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-number-windows
func (s *BootServiceImpl) GetPlesk(ServerNumber string) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/plesk", ServerNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Plesk, resp, err
}

// ActivatePlesk Activate Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-plesk
func (s *BootServiceImpl) ActivatePlesk(req *ActivateWindowsRequest) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/plesk", req.ServerNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Plesk, resp, err
}

// DeactivatePlesk Deactivate Plesk installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-plesk
func (s *BootServiceImpl) DeactivatePlesk(ServerNumber string) (*Plesk, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/plesk", ServerNumber)

	data := dataPlesk{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, nil)
	return data.Plesk, resp, err
}

// GetCPanel Query boot options for the cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#get-boot-server-ip-cpanel
func (s *BootServiceImpl) GetCPanel(ServerNumber string) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/cpanel", ServerNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Cpanel, resp, err
}

// ActivateCPanel Activate cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#post-boot-server-number-cpanel
func (s *BootServiceImpl) ActivateCPanel(req *ActivateCPanelRequest) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/cpanel", req.ServerNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Cpanel, resp, err
}

// DeactivateCPanel Deactivate cPanel installation
// See: https://robot.your-server.de/doc/webservice/en.html#delete-boot-server-number-cpanel
func (s *BootServiceImpl) DeactivateCPanel(ServerNumber string) (*Cpanel, *http.Response, error) {
	path := fmt.Sprintf("/boot/%s/cpanel", ServerNumber)

	data := dataCpanel{}
	resp, err := s.client.Call(http.MethodDelete, path, nil, nil)
	return data.Cpanel, resp, err
}
