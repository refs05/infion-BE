package roles

import (
	rolesUsecase "infion-BE/businesses/roles"
	"time"
)

type Roles struct {
	ID           	int
	Name      		string
	CreatedAt    	time.Time
	UpdatedAt   	time.Time
}

func fromDomain(domain *rolesUsecase.Domain) *Roles {
	return &Roles{
		ID:           	domain.ID,
		Name:      		domain.Name,
	}
}

func (rec *Roles) toDomain() rolesUsecase.Domain {
	return rolesUsecase.Domain{
		ID:           	rec.ID,
		Name:      		rec.Name,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}