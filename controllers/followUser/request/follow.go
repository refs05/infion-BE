package request

import "infion-BE/businesses/followUser"

type FollowUser struct{
	FollowerId 	int `json:"follower_id"`
	UserId 	 int `json:"user_id "`
}

func (follow *FollowUser)ToDomain() *followUser.DomainFollowUser{
	return &followUser.DomainFollowUser{
		FollowerId: follow.FollowerId,
		UserId 	: follow.UserId,
	}
}