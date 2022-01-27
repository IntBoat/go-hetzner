package hetzner

type bootMeta struct {
	ServerIP      string `json:"server_ip"`
	ServerIpv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
	Active        bool   `json:"active"`
}

// Rescue represents Rescue config
type Rescue struct {
	bootMeta
	Os            []string         `json:"os"`
	Arch          []int            `json:"arch"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

// Linux represents Linux boot config
type Linux struct {
	bootMeta
	Dist          []string         `json:"dist"`
	Arch          []int            `json:"arch"`
	Lang          []string         `json:"lang"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

// Vnc represents Vnc boot config
type Vnc struct {
	bootMeta
	Dist     []string `json:"dist"`
	Arch     []int    `json:"arch"`
	Lang     []string `json:"lang"`
	Password *string  `json:"password"`
}

// Windows represents Windows boot config
type Windows struct {
	bootMeta
	Dist     []string `json:"dist"`
	Lang     []string `json:"lang"`
	Password *string  `json:"password"`
}

// Plesk represents Plesk boot config
type Plesk struct {
	bootMeta
	Dist     []string `json:"dist"`
	Arch     []int    `json:"arch"`
	Lang     []string `json:"lang"`
	Password *string  `json:"password"`
	Hostname *string  `json:"hostname"`
}

// Cpanel represents Cpanel boot config
type Cpanel struct {
	bootMeta
	Dist     []string `json:"dist"`
	Arch     []int    `json:"arch"`
	Lang     []string `json:"lang"`
	Password *string  `json:"password"`
	Hostname *string  `json:"hostname"`
}

// Boot represents ALL boot config
type Boot struct {
	Rescue  *Rescue  `json:"rescue"`
	Linux   *Linux   `json:"linux"`
	Vnc     *Vnc     `json:"vnc"`
	Windows *Windows `json:"windows"`
	Plesk   *Plesk   `json:"plesk"`
	Cpanel  *Cpanel  `json:"cpanel"`
}

type requestMeta struct {
	ServerNumber int    // Server ID
	Dist         string `url:"dist"` // Distribution
	Arch         int    `url:"arch"` // Architecture (optional, default: 64)
	Lang         string `url:"lang"` // Language
}

// ActivateRescueRequest represents ActivateRescue request parameter
type ActivateRescueRequest struct {
	requestMeta
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

// ActivateLinuxRequest represents ActivateLinux request parameter
type ActivateLinuxRequest struct {
	requestMeta
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

// ActivateVncRequest represents ActivateVnc request parameter
type ActivateVncRequest struct {
	requestMeta
}

// ActivateWindowsRequest represents ActivateWindows request parameter
type ActivateWindowsRequest struct {
	requestMeta
}

// ActivatePleskRequest represents ActivatePlesk request parameter
type ActivatePleskRequest struct {
	requestMeta
	Hostname *string `url:"hostname"` // Hostname
}

// ActivateCPanelRequest represents ActivateCPanel request parameter
type ActivateCPanelRequest struct {
	requestMeta
	Hostname *string `url:"hostname"` // Hostname
}

type dataLinux struct {
	Linux *Linux `json:"linux"`
}

type dataVnc struct {
	Vnc *Vnc `json:"vnc"`
}

type dataWindows struct {
	Windows *Windows `json:"windows"`
}

type dataPlesk struct {
	Plesk *Plesk `json:"plesk"`
}

type dataCpanel struct {
	Cpanel *Cpanel `json:"cpanel"`
}

type dataBoot struct {
	Boot *Boot `json:"boot"`
}

type dataRescue struct {
	Rescue *Rescue `json:"rescue"`
}
