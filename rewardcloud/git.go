package rewardcloud

type Git struct {
	ID                 int            `json:"id"`
	UUID               string         `json:"uuid"`
	Repo               string         `json:"repo"`
	Username           string         `json:"username"`
	Password           string         `json:"password"`
	SSHPrivateKey      string         `json:"sshPrivateKey"`
	SSHPrivateKeyInput string         `json:"sshPrivateKeyInput"`
	GitType            GitType        `json:"gitType"`
	CredentialType     CredentialType `json:"credentialType"`
}

type GitType struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CredentialType struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
