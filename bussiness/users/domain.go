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
	LikeCount		int
	FollowerCount	int
	CretedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Token    		string

}


type UserUseCaseInterface interface {
	Login(domain DomainUser, ctx context.Context )(DomainUser,error)
	CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error)
}

type UserRepoInterface interface{
	Login(domain DomainUser, ctx context.Context )(DomainUser,error)
	GetUsername(domain DomainUser,ctx context.Context)(DomainUser,error)
	CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error)
}