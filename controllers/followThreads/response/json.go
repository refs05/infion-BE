package response

import (
	"infion-BE/businesses/followThreads"
	"time"
)

type FollowThreads struct {
	ID          	int `json:"id"`
	ThreadID		int	`json:"thread_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain followThreads.Domain) FollowThreads {
	return FollowThreads{
		ID:           	domain.ID,
		ThreadID:		domain.ThreadID,
		UserID: 		domain.UserID,
		Status: 		domain.Status,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}