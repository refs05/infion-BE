package request

import "infion-BE/businesses/users"

type Users struct {
	Id            uint           `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	RoleId        int            `json:"role_id"`
	UrlImg        string         `json:"url_img"`
	CommentCount  int            `json:"comment_count"`
	LikeCount     int            `json:"like_count"`
	FollowerCount int            `json:"follower_count"`
	Token         string         `json:"token"`
}

func (user *Users) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Id:				user.Id,
		Username: 		user.Username,
		Email:  		user.Email,
		Password:   	user.Password,
		RoleId:  		user.RoleId,
		UrlImg: 		user.UrlImg,
		CommentCount: 	user.CommentCount,
		LikeCount: 		user.LikeCount,
		FollowerCount: 	user.FollowerCount,	
		Token:			user.Token, 	
	}
}