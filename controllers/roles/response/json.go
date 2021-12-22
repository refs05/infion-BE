package response

import (
	"infion-BE/businesses/roles"
	"time"
)

type Roles struct {
	ID           	int `json:"id"`
	Name      		string `json:"name"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain roles.Domain) Roles {
	return Roles{
		ID:           	domain.ID,
		Name:      	domain.Name,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}