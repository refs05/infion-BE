package response

import (
	"infion-BE/businesses/likeComments"
	"time"
)

type LikeComments struct {
	ID          	int `json:"id"`
	CommentID		int	`json:"comment_id"`
	UserID			int	`json:"user_id"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain likeComments.Domain) LikeComments {
	return LikeComments{
		ID:           	domain.ID,
		CommentID:		domain.CommentID,
		UserID: 		domain.UserID,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}