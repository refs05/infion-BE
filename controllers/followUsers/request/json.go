package request

import "infion-BE/businesses/followUsers"

type FollowUsers struct {
	ID          	int `json:"id"`
	FollowedID		int	`json:"followed_id"`
	FollowerID		int	`json:"follower_id"`
	Status			bool `json:"status"`
}

func (req *FollowUsers) ToDomain() *followUsers.Domain {
	return &followUsers.Domain{
		ID:				req.ID,
		FollowedID:		req.FollowedID,
		FollowerID: 	req.FollowerID,
		Status: 		req.Status,
	}
}