package request

import "infion-BE/businesses/roles"

type Roles struct {
	ID           	int `json:"id"`
	Name      		string `json:"name"`
}

func (req *Roles) ToDomain() *roles.Domain {
	return &roles.Domain{
		ID:				req.ID,
		Name:      		req.Name,
	}
}