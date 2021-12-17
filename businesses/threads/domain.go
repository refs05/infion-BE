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
	LikeCount		int
	CommentCount	int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	// GetByTitle(ctx context.Context, threadsTitle string) (Domain, error)
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// Update(ctx context.Context, threadsDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	// GetByTitle(ctx context.Context, threadsTitle string) (Domain, error)
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// Update(ctx context.Context, threadsDomain *Domain) (Domain, error)
}