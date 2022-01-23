package response

import "infion-BE/businesses/followUser"

type FollowUserResponse struct {
	Id          uint `json:"id"`
	FollowerId  int `json:"follower_id"`
	UserId 	 int `json:"user_id"`
}

func FromDomain(domain followUser.DomainFollowUser) FollowUserResponse{
	return FollowUserResponse{
		Id: domain.Id,
		FollowerId: domain.FollowerId,
		UserId 	: domain.UserId ,
	}
}