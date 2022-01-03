package comments

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	ThreadID		int
	UserID			int
	User			string
	UrlImg			string
	Comment   		string
	LikeCount		int
	ReplyCount		int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, commentsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, commentsId int) (Domain, error)
	GetComments(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, commentsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, commentsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, commentsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, commentsId int) (Domain, error)
	GetComments(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, commentsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, commentsDomain *Domain) (Domain, error)
}