package threads

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	Title        	string
	Img				string
	Content      	string
	Category   		string
	UserID			int
	User			string
	LikeCount		int
	CommentCount	int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	Update(ctx context.Context, threadsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, threadsDomain *Domain) (*Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	Update(ctx context.Context, threadsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
}