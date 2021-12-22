package roles

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	Name			string
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, rolesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, rolesId int) (Domain, error)
	Update(ctx context.Context, rolesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, rolesDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, rolesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, rolesId int) (Domain, error)
	Update(ctx context.Context, rolesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, rolesDomain *Domain) (Domain, error)
}