package schema

import "time"

// Project defines the schema of a project.
type Project struct {
	ID                    int             `json:"id"`
	UUID                  string          `json:"uuid"`
	Name                  string          `json:"name"`
	IsActive              bool            `json:"isActive"`
	CPU                   int             `json:"cpu"`
	Memory                int             `json:"memory"`
	Storage               int             `json:"storage"`
	Code                  string          `json:"code"`
	Color                 string          `json:"color"`
	IsInitProjectSkeleton bool            `json:"isInitProjectSkeleton"`
	Environment           []string        `json:"environment"`
	Team                  string          `json:"team"`
	Git                   Git             `json:"git"`
	ServiceAccount        ServiceAccount  `json:"serviceAccount"`
	State                 string          `json:"state"`
	ProjectEnvVar         []ProjectEnvVar `json:"projectEnvVar"`
	ProjectTypeVersion    string          `json:"projectTypeVersion"`
	ComponentVersion      []string        `json:"componentVersion"`
	CreatedAt             time.Time       `json:"createdAt"`
	UpdatedAt             time.Time       `json:"updatedAt"`
}

type ProjectEnvVar struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	IsEncrypted bool   `json:"isEncrypted"`
	EnvVarType  string `json:"envVarType"`
}

// ProjectListResponse defines the schema of the response when
// listing projects.
type ProjectListResponse []Project

// // ProjectGetResponse defines the schema of the response when
// // retrieving a single project.
// type ProjectGetResponse *Project
