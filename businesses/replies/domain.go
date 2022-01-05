package replies

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	CommentID int
	UserID    int
	User      string
	UrlImg    string
	Reply     string
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, repliesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, repliesId int) (Domain, error)
	GetReplies(ctx context.Context) ([]Domain, error)
	GetRepliesByCommentID(ctx context.Context, commentId int) ([]Domain, error)
	Update(ctx context.Context, repliesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, repliesDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, repliesDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, repliesId int) (Domain, error)
	GetReplies(ctx context.Context) ([]Domain, error)
	GetRepliesByCommentID(ctx context.Context, commentId int) ([]Domain, error)
	Update(ctx context.Context, repliesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, repliesDomain *Domain) (Domain, error)
}
