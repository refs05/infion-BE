package request

import "infion-BE/bussiness/users"

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *CreateUser) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Username: user.Username,
		Email:user.Email,
		Password: user.Password,
	}
}