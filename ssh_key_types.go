package hetzner

// SSHKey contains information about an SSH key
type SSHKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
	Data        string `json:"data"`
}

// SSHKeyCreateRequest Requested information for SSHKeyService.Create
type SSHKeyCreateRequest struct {
	Name string `url:"name"` // SSH key name
	Data string `url:"data"` // SSH key data in OpenSSH or SSH2 format
}

// SSHKeyUpdateRequest Requested information for SSHKeyService.Update
type SSHKeyUpdateRequest struct {
	Fingerprint string // Key fingerprint
	Name        string `url:"name"` // SSH key name
}

type dataKey struct {
	Key *SSHKey `json:"key"`
}
