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
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// GetByID(ctx context.Context, threadsId int) (Domain, error)
	// GetByTitle(ctx context.Context, threadsTitle string) (Domain, error)
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	// Update(ctx context.Context, threadsDomain *Domain) (*Domain, error)
}

type Repository interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	// GetByID(ctx context.Context, threadsId int) (Domain, error)
	// GetByTitle(ctx context.Context, threadsTitle string) (Domain, error)
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	// Update(ctx context.Context, threadsDomain *Domain) (Domain, error)
}