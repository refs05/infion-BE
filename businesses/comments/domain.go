package comments

import (
	"context"
	"infion-BE/businesses/replies"
	"time"
)

type Domain struct {
	ID           	int
	ThreadID		int
	UserID			int
	User			string
	UrlImg			string
	Comment   		string
	Replies			[]replies.Domain
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
	GetCommentsByThreadID(ctx context.Context, threadId int) ([]Domain, error)
	Update(ctx context.Context, commentsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, commentsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, commentsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, commentsId int) (Domain, error)
	GetComments(ctx context.Context) ([]Domain, error)
	GetCommentsByThreadID(ctx context.Context, threadId int) ([]Domain, error)
	Update(ctx context.Context, commentsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, commentsDomain *Domain) (Domain, error)
	CountByThreadID(ctx context.Context,id int) (int, error)
}