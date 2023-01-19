package rewardcloud

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/rewardenv/rewardcloud-go/rewardcloud/schema"
)

// Project represents a project in the Reward Cloud.
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

// ProjectClient is a client for the project API.
type ProjectClient struct {
	client *Client
}

// GetByID retrieves a server type by its ID. If the server type does not exist, nil is returned.
func (c *ProjectClient) GetByID(ctx context.Context, id int) (*Project, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/projects/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Project
	resp, err := c.client.Do(req, &body)
	if err != nil {
		if IsError(err, ErrorCodeNotFound) {
			return nil, resp, nil
		}
		return nil, nil, err
	}
	return ProjectFromSchema(body), resp, nil
}

// GetByName retrieves a server type by its name. If the server type does not exist, nil is returned.
func (c *ProjectClient) GetByName(ctx context.Context, name string) (*Project, *Response, error) {
	if name == "" {
		return nil, nil, nil
	}
	projects, response, err := c.List(ctx, ProjectListOpts{})
	if len(projects) == 0 {
		return nil, response, err
	}

	return projects[0], response, err
}

// Get retrieves a server type by its ID if the input can be parsed as an integer, otherwise it
// retrieves a server type by its name. If the server type does not exist, nil is returned.
func (c *ProjectClient) Get(ctx context.Context, id int) (*Project, *Response, error) {
	return c.GetByID(ctx, int(id))
}

// ProjectListOpts specifies options for listing server types.
type ProjectListOpts struct {
	ListOpts
	Team string `json:"team"`
	// Name string
	Sort []string
}

func (l ProjectListOpts) values() url.Values {
	vals := l.ListOpts.values()
	if l.Team != "" {
		vals.Add("team", l.Team)
	}
	for _, sort := range l.Sort {
		vals.Add("sort", sort)
	}
	return vals
}

// List returns a list of server types for a specific page.
//
// Please note that filters specified in opts are not taken into account
// when their value corresponds to their zero value or when they are empty.
func (c *ProjectClient) List(ctx context.Context, opts ProjectListOpts) ([]*Project, *Response, error) {
	path := "/projects?" + opts.values().Encode()
	req, err := c.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.ProjectListResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	projects := make([]*Project, 0, len(body))
	for _, s := range body {
		projects = append(projects, ProjectFromSchema(s))
	}
	return projects, resp, nil
}

// All returns all server types.
func (c *ProjectClient) All(ctx context.Context) ([]*Project, error) {
	allProjects := []*Project{}

	opts := ProjectListOpts{}
	opts.PerPage = 50

	err := c.client.all(func(page int) (*Response, error) {
		opts.Page = page
		projects, resp, err := c.List(ctx, opts)
		if err != nil {
			return resp, err
		}
		allProjects = append(allProjects, projects...)
		return resp, nil
	})
	if err != nil {
		return nil, err
	}

	return allProjects, nil
}
