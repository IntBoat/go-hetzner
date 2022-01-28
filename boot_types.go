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
	Os            interface{}      `json:"os"`
	Arch          interface{}      `json:"arch"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

// Linux represents Linux boot config
type Linux struct {
	bootMeta
	Dist          interface{}      `json:"dist"`
	Arch          interface{}      `json:"arch"`
	Lang          interface{}      `json:"lang"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

// Vnc represents Vnc boot config
type Vnc struct {
	bootMeta
	Dist     interface{} `json:"dist"`
	Arch     interface{} `json:"arch"`
	Lang     interface{} `json:"lang"`
	Password *string     `json:"password"`
}

// Windows represents Windows boot config
type Windows struct {
	bootMeta
	Dist     interface{} `json:"dist"`
	Lang     interface{} `json:"lang"`
	Password *string     `json:"password"`
}

// Plesk represents Plesk boot config
type Plesk struct {
	bootMeta
	Dist     interface{} `json:"dist"`
	Arch     interface{} `json:"arch"`
	Lang     interface{} `json:"lang"`
	Password *string     `json:"password"`
	Hostname *string     `json:"hostname"`
}

// Cpanel represents Cpanel boot config
type Cpanel struct {
	bootMeta
	Dist     interface{} `json:"dist"`
	Arch     interface{} `json:"arch"`
	Lang     interface{} `json:"lang"`
	Password *string     `json:"password"`
	Hostname *string     `json:"hostname"`
}

// ActivateRescueRequest represents ActivateRescue request parameter
type ActivateRescueRequest struct {
	ServerNumber  int       // Server ID
	OS            string    `url:"os"`                      // Operating System
	Arch          int       `url:"arch"`                    // Architecture (optional, default: 64)
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

// ActivateLinuxRequest represents ActivateLinux request parameter
type ActivateLinuxRequest struct {
	ServerNumber  int       // Server ID
	Dist          string    `url:"dist"`                    // Distribution
	Arch          int       `url:"arch"`                    // Architecture (optional, default: 64)
	Lang          string    `url:"lang"`                    // Language
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

// ActivateVncRequest represents ActivateVnc request parameter
type ActivateVncRequest struct {
	ServerNumber int    // Server ID
	Dist         string `url:"dist"` // Distribution
	Arch         int    `url:"arch"` // Architecture (optional, default: 64)
	Lang         string `url:"lang"` // Language
}

// ActivateWindowsRequest represents ActivateWindows request parameter
type ActivateWindowsRequest struct {
	ServerNumber int    // Server ID
	Lang         string `url:"lang"` // Language
}

// ActivatePleskRequest represents ActivatePlesk request parameter
type ActivatePleskRequest struct {
	ServerNumber int     // Server ID
	Dist         string  `url:"dist"`     // Distribution
	Arch         int     `url:"arch"`     // Architecture (optional, default: 64)
	Lang         string  `url:"lang"`     // Language
	Hostname     *string `url:"hostname"` // Hostname
}

// ActivateCPanelRequest represents ActivateCPanel request parameter
type ActivateCPanelRequest struct {
	ServerNumber int     // Server ID
	Dist         string  `url:"dist"`     // Distribution
	Arch         int     `url:"arch"`     // Architecture (optional, default: 64)
	Lang         string  `url:"lang"`     // Language
	Hostname     *string `url:"hostname"` // Hostname
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
