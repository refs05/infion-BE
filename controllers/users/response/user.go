package response

import (
	"infion-BE/businesses/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id            uint           `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	RoleId        int            `json:"role_id"`
	UrlImg        string         `json:"url_img"`
	CommentCount  int            `json:"comment_count"`
	LikeCount     int            `json:"like_count"`
	FollowerCount int            `json:"follower_count"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Token         string         `json:"token"`
}

func FromDomain(domain users.DomainUser) UserResponse{
	return UserResponse{
		Id:				domain.Id,
		Username: 		domain.Username,
		Email:  		domain.Email,
		Password:   	domain.Password,
		RoleId:  		domain.RoleId,
		UrlImg: 		domain.UrlImg,
		CommentCount: 	domain.CommentCount,
		LikeCount: 		domain.LikeCount,
		FollowerCount: 	domain.FollowerCount,		
		CreatedAt:		 domain.CreatedAt,
		UpdatedAt:		 domain.UpdatedAt,
		DeletedAt:		 domain.DeletedAt,
		Token:			 domain.Token, 

	}
}

func NewResponseArray(domainUsers []users.DomainUser) []UserResponse {
	var resp []UserResponse

	for _, value := range domainUsers {
		resp = append(resp, FromDomain(value))
	}

	return resp
}