package request

import "infion-BE/businesses/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Email: user.Email,
		Password: user.Password,
	}
}