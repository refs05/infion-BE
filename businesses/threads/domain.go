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
	UrlImg			string
	LikeCount		int
	CommentCount	int
	FollowerCount	int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	GetThreadsBySort(ctx context.Context, sort string) ([]Domain, error)
	GetThreadsByCategory(ctx context.Context, category string) ([]Domain, error)
	GetThreadsBySortCategory(ctx context.Context, sort string, category string) ([]Domain, error)
	GetThreadsByUserID(ctx context.Context, threadsID int) ([]Domain, error)
	Update(ctx context.Context, threadsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, threadsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, threadsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, threadsId int) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	GetThreadsBySort(ctx context.Context, sort string) ([]Domain, error)
	GetThreadsByCategory(ctx context.Context, category string) ([]Domain, error)
	GetThreadsBySortCategory(ctx context.Context, sort string, category string) ([]Domain, error)
	GetThreadsByUserID(ctx context.Context, userID int) ([]Domain, error)
	Update(ctx context.Context, threadsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, threadsDomain *Domain) (Domain, error)
	CountByUserID(ctx context.Context,id uint) (int, error)
	GetThreadLikeCountByUserID(ctx context.Context, userID uint) (int, error)
}