package users

import (
	"context"
	"errors"
	"infion-BE/businesses/users"

	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB)users.Repository{
	return &UserRepository{
		db:gormDb,
	}
}


func (repo *UserRepository)Login(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)
	err:= repo.db.Where("email = ? ",userDb.Email).First(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

func (repo *UserRepository) GetUsername(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	data := FromDomain(domain)
	err := repo.db.Where("username = ? OR email = ?",data.Username,data.Email).First(&data).Error
	
	errors.Is(err, gorm.ErrRecordNotFound)
	if err ==  gorm.ErrRecordNotFound{
		return data.ToDomain(),gorm.ErrRecordNotFound
	}else if err == nil{
		
		return users.DomainUser{},err
	}else{
		return users.DomainUser{},err
	}
	
	

}

func (repo *UserRepository) CreateNewUser(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Create(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

func (repo *UserRepository) FindById(userId int,ctx context.Context)(users.DomainUser,error){
	user := User{}
	err := repo.db.First(&user, userId).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return user.ToDomain(),nil
}

// func (repo *UserRepository) GetEmail(ctx context.Context,email string)(users.DomainUser,error){
// 	data := User{}

// 	err := repo.db.Where("email = ?",email).First(&data).Error

// 	if err != nil{

// 	}
// }

func (repo *UserRepository) GetLeaderboard(ctx context.Context)([]users.DomainUser, error){
	var recordUser []User
	
	result := repo.db.Unscoped().Order("comment_count desc").Order("like_count desc").Order("follower_count desc").Find(&recordUser)
	if result.Error != nil {
		return []users.DomainUser{}, result.Error
	}

	return ToDomainArray(recordUser), nil
}