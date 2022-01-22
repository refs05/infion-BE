package response

import (
	"infion-BE/businesses/likeComments"
	"time"
)

type LikeComments struct {
	ID          	int `json:"id"`
	CommentID		int	`json:"comment_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain likeComments.Domain) LikeComments {
	return LikeComments{
		ID:           	domain.ID,
		CommentID:		domain.CommentID,
		UserID: 		domain.UserID,
		Status: 		domain.Status,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}