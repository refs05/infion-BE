package followUsers

import (
	"context"
	"infion-BE/businesses/followUsers"

	"gorm.io/gorm"
)

type mysqlFollowUsersRepository struct {
	Conn *gorm.DB
}

func NewFollowUsersRepository(conn *gorm.DB) followUsers.Repository {
	return &mysqlFollowUsersRepository{
		Conn: conn,
	}
}

func (nr *mysqlFollowUsersRepository) Store(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	rec := fromDomain(followUsersDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return followUsers.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowUsersRepository) GetByID(ctx context.Context, followUsersId int) (followUsers.Domain, error) {
	rec := FollowUsers{}
	err := nr.Conn.Where("id = ?", followUsersId).First(&rec).Error
	if err != nil {
		return followUsers.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlFollowUsersRepository) Update(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	rec := fromDomain(followUsersDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return followUsers.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowUsersRepository) Delete(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	rec := fromDomain(followUsersDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return followUsers.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowUsersRepository) CountByFollowedID(ctx context.Context,id uint) (int, error) {
	rec := FollowUsers{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("followed_id = ?", id).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlFollowUsersRepository) GetDuplicate(ctx context.Context, followedID int, followerID int) (followUsers.Domain, error) {
	rec := FollowUsers{}
	err := nr.Conn.Where("followed_id = ? AND follower_id = ?", followedID, followerID).First(&rec).Error
	if err != nil {
		return followUsers.Domain{}, err
	}
	return rec.toDomain(), nil
}