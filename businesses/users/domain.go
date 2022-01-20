package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)
type DomainUser struct{
	Id				uint
	Username		string
	Email 			string
	Password 		string
	RoleId			int
	UrlImg			string
	CommentCount	int
	ThreadCount		int
	LikeCount		int
	FollowerCount	int
	ThreadFollowerCount  int
	ThreadFollowingCount int
	Rank            int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Token    		string

}


type UseCase interface {
	Login(domain DomainUser, ctx context.Context )(DomainUser,error)
	CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error)
	FindById(userId int,ctx context.Context)(DomainUser,error)
	Update(userDomain *DomainUser, ctx context.Context, ) (*DomainUser, error)
	GetLeaderboard(ctx context.Context) ([]DomainUser, error)
}

type Repository interface{
	Login(domain DomainUser, ctx context.Context )(DomainUser,error)
	GetUsername(domain DomainUser,ctx context.Context)(DomainUser,error)
	CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error)
	FindById(userId int,ctx context.Context)(DomainUser,error)
	Update(userDomain *DomainUser, ctx context.Context) (DomainUser, error)
	GetLeaderboard(ctx context.Context) ([]DomainUser, error)
}