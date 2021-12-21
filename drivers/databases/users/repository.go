package users

import (
	"context"
	"errors"
	"infion-BE/bussiness/users"

	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB)users.UserRepoInterface{
	return &UserRepository{
		db:gormDb,
	}
}


func (repo *UserRepository)Login(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)
	err:= repo.db.Where("email = ? AND password = ?",userDb.Email,userDb.Password).First(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

func (repo *UserRepository) GetUsername(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	data := FromDomain(domain)
	res := repo.db.Where("username = ? AND email = ?",data.Username,data.Email).First(&data)
	if res != nil{
		return users.DomainUser{},errors.New("username or email already exists")
	}
	return data.ToDomain(),nil

}

func (repo *UserRepository) CreateNewUser(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Create(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}
// func (repo *UserRepository) GetEmail(ctx context.Context,email string)(users.DomainUser,error){
// 	data := User{}

// 	err := repo.db.Where("email = ?",email).First(&data).Error

// 	if err != nil{

// 	}
// }