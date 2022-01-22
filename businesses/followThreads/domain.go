package followThreads

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	ThreadID		int
	UserID			int
	Status			bool
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, followThreadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, followThreadsId int) (Domain, error)
	Update(ctx context.Context, followThreadsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, followThreadsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, followThreadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, followThreadsId int) (Domain, error)
	GetFollowThreadsByThreadID(ctx context.Context, threadID int) ([]Domain, error)
	Update(ctx context.Context, followThreadsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, followThreadsDomain *Domain) (Domain, error)
	CountByThreadID(ctx context.Context,id int) (int, error)
	CountByUserID(ctx context.Context,id uint) (int, error)
	GetDuplicate(ctx context.Context, threadID int, userID int) (Domain, error)
}