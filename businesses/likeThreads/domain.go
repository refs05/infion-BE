package likeThreads

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	ThreadID		int
	UserID			int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, likeThreadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeThreadsId int) (Domain, error)
	Update(ctx context.Context, likeThreadsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, likeThreadsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, likeThreadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeThreadsId int) (Domain, error)
	GetLikeThreadsByThreadID(ctx context.Context, threadID int) ([]Domain, error)
	Update(ctx context.Context, likeThreadsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, likeThreadsDomain *Domain) (Domain, error)
	CountByThreadID(ctx context.Context,id int) (int, error)
	GetDuplicate(ctx context.Context, threadID int, userID int) (Domain, error)
}