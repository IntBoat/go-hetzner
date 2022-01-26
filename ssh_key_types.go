package hetzner

type SSHKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
	Data        string `json:"data"`
}

type SSHKeyCreateRequest struct {
	Name string `url:"name"` // SSH key name
	Data string `url:"data"` // SSH key data in OpenSSH or SSH2 format
}

type SSHKeyUpdateRequest struct {
	Fingerprint string // Key fingerprint
	Name        string `url:"name"` // SSH key name
}

type dataKey struct {
	Key *SSHKey `json:"key"`
}
