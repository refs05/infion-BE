package followUsers

import (
	"context"
	"time"
)

type Domain struct {
	ID				int
	FollowedID		int
	FollowerID		int
	CreatedAt		time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, followUsersDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, followUsersId int) (Domain, error)
	Update(ctx context.Context, followUsersDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, followUsersDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, followUsersDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, followUsersId int) (Domain, error)
	Update(ctx context.Context, followUsersDomain *Domain) (Domain, error)
	Delete(ctx context.Context, followUsersDomain *Domain) (Domain, error)
	CountByFollowedID(ctx context.Context, id int) (int, error)
}