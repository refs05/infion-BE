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
	RoleId			int `gorm:"default:1"`
	UrlImg			string
	CommentCount	int
	ThreadCount		int
	LikeCount		int
	FollowerCount	int
	ThreadFollowerCount  int
	ThreadFollowingCount int
	Rank            int
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
		ThreadCount:	domain.ThreadCount,
		LikeCount:		domain.LikeCount,
		FollowerCount:	domain.FollowerCount,
		ThreadFollowerCount:  domain.ThreadFollowerCount,
		ThreadFollowingCount: domain.ThreadFollowingCount,
		Rank:           domain.Rank,
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
		ThreadCount: 		user.ThreadCount,
		LikeCount:			user.LikeCount,
		FollowerCount:		user.FollowerCount,
		ThreadFollowerCount:  user.ThreadFollowerCount,
		ThreadFollowingCount: user.ThreadFollowingCount,
		Rank:               user.Rank,
		Password:  			user.Password,
		CreatedAt:			user.CreatedAt,
		UpdatedAt:			user.UpdatedAt,
		DeletedAt: 			user.DeletedAt,
	}
}

func ToDomainArray(modelThreads []User) []users.DomainUser {
	var response []users.DomainUser

	for _, val := range modelThreads{
		response = append(response, val.ToDomain())
	}
	return response
}