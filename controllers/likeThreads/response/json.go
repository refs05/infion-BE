package response

import (
	"infion-BE/businesses/likeThreads"
	"time"
)

type LikeThreads struct {
	ID          	int `json:"id"`
	ThreadID		int	`json:"thread_id"`
	UserID			int	`json:"user_id"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain likeThreads.Domain) LikeThreads {
	return LikeThreads{
		ID:           	domain.ID,
		ThreadID:		domain.ThreadID,
		UserID: 		domain.UserID,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}