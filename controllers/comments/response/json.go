package response

import (
	"infion-BE/businesses/comments"
	"infion-BE/businesses/replies"
	"time"
)

type Comments struct {
	ID           	int `json:"id"`
	ThreadID		int `json:"thread_id"`
	UserID			int `json:"user_id"`
	User			string `json:"username"`
	UrlImg			string `json:"url_img"`
	Comment			string `json:"comment"`
	Replies			[]replies.Domain `json:"replies"`
	LikeCount		int `json:"like_count"`
	ReplyCount		int `json:"reply_count"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain comments.Domain) Comments {
	return Comments{
		ID:           	domain.ID,
		ThreadID: 		domain.ThreadID,
		UserID: 		domain.UserID,
		User: 			domain.User,
		UrlImg:			domain.UrlImg,
		Comment: 		domain.Comment,
		Replies:		domain.Replies,
		LikeCount:  	domain.LikeCount,
		ReplyCount: 	domain.ReplyCount,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}

func NewResponseArray(domainComments []comments.Domain) []Comments {
	var resp []Comments

	for _, value := range domainComments {
		resp = append(resp, FromDomain(value))
	}

	return resp
}