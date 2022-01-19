package followUser

import (
	"context"
	"infion-BE/businesses/followUser"

	"gorm.io/gorm"
)

type FollowUserRepository struct {
	db *gorm.DB
}
func NewFollowUserRepository(gormDb *gorm.DB)followUser.Repository{
	return &FollowUserRepository{
		db:gormDb,
	}
}

func (repo *FollowUserRepository) Follow(domain followUser.DomainFollowUser,ctx context.Context)(followUser.DomainFollowUser,error){
	followDb := FromDomain(domain)
	err := repo.db.Create(&followDb).Error

	if err != nil{
		return followUser.DomainFollowUser{},err
	}
	return followDb.ToDomain(),nil
}
