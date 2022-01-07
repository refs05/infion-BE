package comments

import (
	commentsUsecase "infion-BE/businesses/comments"
	repliesUsecase "infion-BE/businesses/replies"
	"infion-BE/drivers/databases/replies"
	"infion-BE/drivers/databases/threads"
	"infion-BE/drivers/databases/users"
	"time"
)

type Comments struct {
	ID           	int `gorm:"primaryKey"`
	ThreadID		int
	Thread			threads.Threads
	UserID			int
	User			users.User
	UrlImg			string
	Comment			string
	Replies			[]replies.Replies `gorm:"foreignKey:CommentID;references:ID"`
	LikeCount		int
	ReplyCount		int
	CreatedAt    	time.Time
	UpdatedAt   	time.Time
}

func fromDomain(domain *commentsUsecase.Domain) *Comments {
	return &Comments{
		ID:           	domain.ID,
		ThreadID: 		domain.ThreadID,
		UserID: 		domain.UserID,
		Comment:		domain.Comment,
		LikeCount:  	domain.LikeCount,
		ReplyCount: 	domain.ReplyCount,
	}
}

func (rec *Comments) toDomain() commentsUsecase.Domain {
	rep := ToDomainArray2(rec.Replies)
	return commentsUsecase.Domain{
		ID:           	rec.ID,
		ThreadID: 		rec.ThreadID,
		UserID: 		rec.UserID,
		User:			rec.User.Username,
		UrlImg: 		rec.User.UrlImg,
		Comment: 		rec.Comment,
		Replies: 		rep,
		LikeCount:  	rec.LikeCount,
		ReplyCount: 	rec.ReplyCount,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}

func ToDomainArray(modelComments []Comments) []commentsUsecase.Domain {
	var response []commentsUsecase.Domain

	for _, val := range modelComments{
		response = append(response, val.toDomain())
	}
	return response
}

func toDomain2(rec replies.Replies) repliesUsecase.Domain {
	return repliesUsecase.Domain{
		ID:        rec.ID,
		CommentID: rec.CommentID,
		UserID:    rec.UserID,
		User:      rec.User.Username,
		UrlImg:    rec.User.UrlImg,
		Reply:     rec.Reply,
		LikeCount: rec.LikeCount,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainArray2(modelReplies []replies.Replies) []repliesUsecase.Domain {
	var response []repliesUsecase.Domain

	for _, val := range modelReplies{
		response = append(response, toDomain2(val))
	}
	return response
}