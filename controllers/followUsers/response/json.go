package response

import (
	"infion-BE/businesses/followUsers"
	"time"
)

type FollowUsers struct {
	ID          	int `json:"id"`
	FollowedID		int	`json:"followed_id"`
	FollowerID		int	`json:"follower_id"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain followUsers.Domain) FollowUsers {
	return FollowUsers{
		ID:           	domain.ID,
		FollowedID:		domain.FollowedID,
		FollowerID: 	domain.FollowerID,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}