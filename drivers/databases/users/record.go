package users

import (
	"infion-BE/businesses/users"
	"time"

	"gorm.io/gorm"
)
type User struct {
	Id        		uint `gorm:"primaryKey"`
	Email     		string
	Username      	string
	RoleId			int
	UrlImg			string
	CommentCount	int
	LikeCount		int
	FollowerCount	int
	Password  		string
	CreatedAt  		time.Time `gorm:"<-:create"`
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	
}

func FromDomain(domain users.DomainUser) User{
	return User{
		Id: domain.Id,
		Email:     		domain.Email,
		Username:      	domain.Username,
		RoleId:			domain.RoleId,
		UrlImg:			domain.UrlImg,
		CommentCount:	domain.CommentCount,
		LikeCount:		domain.LikeCount,
		FollowerCount:	domain.FollowerCount,
		Password:  		domain.Password,
		CreatedAt:		domain.CreatedAt,
		UpdatedAt:		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}

func (user User)ToDomain() users.DomainUser{
	return users.DomainUser{
		Id: 				user.Id,
		Email:     			user.Email,
		Username:      		user.Username,
		RoleId:				user.RoleId,
		UrlImg:				user.UrlImg,
		CommentCount:		user.CommentCount,
		LikeCount:			user.LikeCount,
		FollowerCount:		user.FollowerCount,
		Password:  			user.Password,
		CreatedAt:			user.CreatedAt,
		UpdatedAt:			user.UpdatedAt,
		DeletedAt: 			user.DeletedAt,
	}
}