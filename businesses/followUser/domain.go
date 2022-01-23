package followUser

import "context"

type DomainFollowUser struct {
	Id				uint
	FollowerId  	int
	UserId 			int
}

type UseCase interface {
	Follow(domain DomainFollowUser, ctx context.Context) (DomainFollowUser, error)
	// Unfollow(domain DomainFollowUser,ctx context.Context)(DomainFollowUser,error)
}

type Repository interface {
	Follow(domain DomainFollowUser, ctx context.Context) (DomainFollowUser, error)
	// Unfollow(domain DomainFollowUser,ctx context.Context)(DomainFollowUser,error)
}