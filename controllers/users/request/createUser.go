package request

import "infion-BE/businesses/users"

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UrlImg   string `json:"url_img"`
}

func (user *CreateUser) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Username: user.Username,
		Email:user.Email,
		Password: user.Password,
		UrlImg: user.UrlImg,
	}
}