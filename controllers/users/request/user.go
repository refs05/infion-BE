package request

import "infion-BE/businesses/users"

type Users struct {
	Id            uint           `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	RoleId        int            `json:"role_id"`
	UrlImg        string         `json:"url_img"`
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
		Token:			user.Token, 	
	}
}