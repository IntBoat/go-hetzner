package hetzner

import "testing"

func TestList(t *testing.T) {
	ssh, _, err := client.SSHKey.List()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Found %d SSH keys", len(ssh))
}

var sshKeyFingerprint string

func TestCreate(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	ssh, _, err := client.SSHKey.Create(
		&SSHKeyCreateRequest{
			Name: "API Testing",
			Data: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8XG/CHPsyY7Ytbdo+OfvoeZm9xATDu84HuvUi/EW+W05ywOz19MY0klzo4vrXQBAwWMl06oGRfroe5ImJWoX7/s4mL2Xe0nXiZ+cMPwD6x0W+xHpQPyZ18apfFEXq3rhxUhssDG++VnjY9RhRlzZadgRm5mTqvWGXHymhQ+1ggnIwfBVgl5cZHMVqKDe1fKEG7X6gVXBM/Crc568a8ADwO152Tq2oUMzTnNobQjAbCFCyJ0qSEOgd3OP80uCfMyzAOQ9MAqrkKJckV3To3FCDO57+lKNU7W7y/dCHKKaTs28TDwDGX3jdPUnO/PAk5rf+5t1Jap+gZ4RagaM0E/Mgt106yngotLTYjin0XricxKDCdqpZlE5Ohf4tqekzdklkiZvuuQ1RjcD7szIW11urAr+l+43eWHYGZJwafTbKDv0c0kV7UuAydVZvsjSXhM9Y0EaYwd0RkD7/NW7lWpSEsvtRDCDTUPAISUmswcipC9jdf479CuM5AcB/Y3Pu5pnORgGTwWdcCp/uoFo2L9UgUtCguL2IHzG/IeH3hFrpXOtsDmWQXgdXdSdIg/9zD+aD2JAanCiNqmF8BnZz8Zx/dskhrqhf2T9ewaeWjd+vQr8cozv3BeKp4OjwGDMINb0PhqMUvTZY2xSuhdz/D3PqoTVpN9Q52hg0/RAhLOp/uQ== Generated By Termius`,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("SSH: %s", ssh.Fingerprint)
	sshKeyFingerprint = ssh.Fingerprint
}

func TestGet(t *testing.T) {
	ssh, _, err := client.SSHKey.Get(sshKeyFingerprint)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("SSH: %s", ssh.Fingerprint)
}

func TestUpdate(t *testing.T) {
	ssh, _, err := client.SSHKey.Update(
		&SSHKeyUpdateRequest{
			Fingerprint: sshKeyFingerprint,
			Name:        "API Testing Rename",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("SSH: %s", ssh.Name)
}

func TestDelete(t *testing.T) {
	resp, err := client.SSHKey.Delete(sshKeyFingerprint)

	if err != nil && resp.StatusCode != 200 {
		t.Error(err)
		return
	}

	if resp.StatusCode != 200 {
		t.Error("Failed to remove SSH Key entry")
	}
}
