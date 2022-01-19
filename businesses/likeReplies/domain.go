package likeReplies

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	ReplyID			int
	UserID			int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, likeRepliesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeRepliesId int) (Domain, error)
	Update(ctx context.Context, likeRepliesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, likeRepliesDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, likeRepliesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeRepliesId int) (Domain, error)
	Update(ctx context.Context, likeRepliesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, likeRepliesDomain *Domain) (Domain, error)
	CountByReplyID(ctx context.Context,id int) (int, error)
}