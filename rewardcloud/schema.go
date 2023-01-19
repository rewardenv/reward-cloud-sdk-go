package rewardcloud

import "github.com/rewardenv/rewardcloud-go/rewardcloud/schema"

// ProjectFromSchema converts a schema.Project to a Project.
func ProjectFromSchema(s schema.Project) *Project {
	return &Project{
		ID:   s.ID,
		Name: s.Name,
	}
}

// PaginationFromSchema converts a schema.MetaPagination to a Pagination.
func PaginationFromSchema(s schema.MetaPagination) Pagination {
	return Pagination{
		Page:         s.Page,
		PerPage:      s.PerPage,
		PreviousPage: s.PreviousPage,
		NextPage:     s.NextPage,
		LastPage:     s.LastPage,
		TotalEntries: s.TotalEntries,
	}
}

// ErrorFromSchema converts a schema.Error to an Error.
func ErrorFromSchema(s schema.Error) Error {
	e := Error{
		Code:    ErrorCode(s.Code),
		Message: s.Message,
	}

	if d, ok := s.Details.(schema.ErrorDetailsInvalidInput); ok {
		details := ErrorDetailsInvalidInput{
			Fields: []ErrorDetailsInvalidInputField{},
		}
		for _, field := range d.Fields {
			details.Fields = append(details.Fields, ErrorDetailsInvalidInputField{
				Name:     field.Name,
				Messages: field.Messages,
			})
		}
		e.Details = details
	}
	return e
}
