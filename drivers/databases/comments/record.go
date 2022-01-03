package comments

import (
	commentsUsecase "infion-BE/businesses/comments"
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
	return commentsUsecase.Domain{
		ID:           	rec.ID,
		ThreadID: 		rec.ThreadID,
		UserID: 		rec.UserID,
		User:			rec.User.Username,
		UrlImg: 		rec.User.UrlImg,
		Comment: 		rec.Comment,
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