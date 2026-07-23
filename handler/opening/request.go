package opening

import (
	"fmt"
	"goportunities/handler"
)

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or invalid.")
	}
	if r.Role == "" {
		return handler.ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return handler.ErrParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return handler.ErrParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return handler.ErrParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return handler.ErrParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return handler.ErrParamIsRequired("salary", "int64")
	}

	return nil
}

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is trythy
	if r.Role != "" || r.Company != "" || r.Location != "" ||
		r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	// if name of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
