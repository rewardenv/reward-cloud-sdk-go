package rewardcloud

type ServiceAccountGit struct {
	IsExternal    bool   `json:"isExternal"`
	URL           string `json:"url"`
	SSHPrivateKey string `json:"sshPrivateKey"`
}

type ServiceAccountRegistry struct {
	URL string `json:"url"`
}

type ServiceAccount struct {
	ServiceAccountGit      ServiceAccountGit      `json:"serviceAccountGit"`
	ServiceAccountRegistry ServiceAccountRegistry `json:"serviceAccountRegistry"`
}
