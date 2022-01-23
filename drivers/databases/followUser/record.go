package followUser

import "infion-BE/businesses/followUser"

type FollowUser struct {
	Id          uint `gorm:"primaryKey"`
	FollowerId  int
	UserId 		 int
}

func FromDomain(domain followUser.DomainFollowUser) FollowUser{
	return FollowUser{
		Id:			 domain.Id,
		FollowerId:  domain.FollowerId,
		UserId 	: domain.UserId ,
	}
}

func (follow FollowUser)ToDomain() followUser.DomainFollowUser{
return followUser.DomainFollowUser{
	Id:				follow.Id,
	FollowerId: 	follow.FollowerId,
	UserId 	:  	follow.UserId,
}
}