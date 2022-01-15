package replies

import (
	"context"
	"time"
)

type Domain struct {
	ID        int       `json:"id"`
	CommentID int       `json:"comment_id"`
	UserID    int       `json:"user_id"`
	User      string    `json:"username"`
	UrlImg    string    `json:"url_img"`
	Reply     string    `json:"reply"`
	LikeCount int       `json:"like_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
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
	CountByCommentID(ctx context.Context,id int) (int, error)
}
