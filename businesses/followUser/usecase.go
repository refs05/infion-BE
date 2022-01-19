package followUser

import (
	"context"
	"time"
)

type FollowUserUseCase struct {
	repo Repository
	ctx  time.Duration
}

func NewFollowUserUseCase(FollowUserRepo Repository,contextTimeout time.Duration) UseCase{
	return &FollowUserUseCase{
		repo: FollowUserRepo,
		ctx: contextTimeout,
	}
}

func (usecase *FollowUserUseCase) Follow(domain DomainFollowUser,ctx context.Context)(DomainFollowUser,error){
	rec,err := usecase.repo.Follow(domain,ctx)
	if err != nil{
		return DomainFollowUser{},err
	}
	return rec,nil
}

// func (UseCase *FollowUserUseCase) Unfollow(domain DomainFollowUser,ctx context.Context)(DomainFollowUser,error){

// }