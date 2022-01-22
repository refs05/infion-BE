package likeComments

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	CommentID		int
	UserID			int
	CreatedAt    	time.Time
	UpdatedAt    	time.Time
	DeletedAt		time.Time
}

type Usecase interface {
	Store(ctx context.Context, likeCommentsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeCommentsId int) (Domain, error)
	Update(ctx context.Context, likeCommentsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, likeCommentsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, likeCommentsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, likeCommentsId int) (Domain, error)
	GetLikeCommentsByCommentID(ctx context.Context, commentID int) ([]Domain, error)
	Update(ctx context.Context, likeCommentsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, likeCommentsDomain *Domain) (Domain, error)
	CountByCommentID(ctx context.Context,id int) (int, error)
	GetDuplicate(ctx context.Context, commentID int, userID int) (Domain, error)
}