package hetzner

type Rescue struct {
	ServerIp      string           `json:"server_ip"`
	ServerIpv6Net string           `json:"server_ipv6_net"`
	ServerNumber  int              `json:"server_number"`
	Os            []string         `json:"os"`
	Arch          []int            `json:"arch"`
	Active        bool             `json:"active"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

type Linux struct {
	ServerIp      string           `json:"server_ip"`
	ServerIpv6Net string           `json:"server_ipv6_net"`
	ServerNumber  int              `json:"server_number"`
	Dist          []string         `json:"dist"`
	Arch          []int            `json:"arch"`
	Lang          []string         `json:"lang"`
	Active        bool             `json:"active"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

type Vnc struct {
	ServerIp      string   `json:"server_ip"`
	ServerIpv6Net string   `json:"server_ipv6_net"`
	ServerNumber  int      `json:"server_number"`
	Dist          []string `json:"dist"`
	Arch          []int    `json:"arch"`
	Lang          []string `json:"lang"`
	Active        bool     `json:"active"`
	Password      *string  `json:"password"`
}

type Windows struct {
	ServerIp      string   `json:"server_ip"`
	ServerIpv6Net string   `json:"server_ipv6_net"`
	ServerNumber  int      `json:"server_number"`
	Dist          []string `json:"dist"`
	Lang          []string `json:"lang"`
	Active        bool     `json:"active"`
	Password      *string  `json:"password"`
}

type Plesk struct {
	ServerIp      string   `json:"server_ip"`
	ServerIpv6Net string   `json:"server_ipv6_net"`
	ServerNumber  int      `json:"server_number"`
	Dist          []string `json:"dist"`
	Arch          []int    `json:"arch"`
	Lang          []string `json:"lang"`
	Active        bool     `json:"active"`
	Password      *string  `json:"password"`
	Hostname      *string  `json:"hostname"`
}

type Cpanel struct {
	ServerIp      string   `json:"server_ip"`
	ServerIpv6Net string   `json:"server_ipv6_net"`
	ServerNumber  int      `json:"server_number"`
	Dist          []string `json:"dist"`
	Arch          []int    `json:"arch"`
	Lang          []string `json:"lang"`
	Active        bool     `json:"active"`
	Password      *string  `json:"password"`
	Hostname      *string  `json:"hostname"`
}

type Boot struct {
	Rescue  *Rescue  `json:"rescue"`
	Linux   *Linux   `json:"linux"`
	Vnc     *Vnc     `json:"vnc"`
	Windows *Windows `json:"windows"`
	Plesk   *Plesk   `json:"plesk"`
	Cpanel  *Cpanel  `json:"cpanel"`
}

type ActiveMeta struct {
	ServerNumber string // Server ID
	Dist         string `url:"dist"` // Distribution
	Arch         int    `url:"arch"` // Architecture (optional, default: 64)
	Lang         string `url:"lang"` // Language
}

type ActivateLinuxRequest struct {
	ActiveMeta
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

type ActivateRescueRequest struct {
	ActiveMeta
	AuthorizedKey *[]string `url:"authorized_key,brackets"` // One or more SSH key fingerprints (optional)
}

type ActivateVncRequest struct {
	ActiveMeta
}

type ActivateWindowsRequest struct {
	ActiveMeta
}

type ActivatePleskRequest struct {
	ActiveMeta
	Hostname *string `url:"hostname"` // Hostname
}

type ActivateCPanelRequest struct {
	ActiveMeta
	hostname *string `url:"hostname"` // Hostname
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
