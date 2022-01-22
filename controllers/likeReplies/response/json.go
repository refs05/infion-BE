package response

import (
	"infion-BE/businesses/likeReplies"
	"time"
)

type LikeReplies struct {
	ID          	int `json:"id"`
	ReplyID			int	`json:"reply_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain likeReplies.Domain) LikeReplies {
	return LikeReplies{
		ID:           	domain.ID,
		ReplyID:		domain.ReplyID,
		UserID: 		domain.UserID,
		Status: 		domain.Status,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}