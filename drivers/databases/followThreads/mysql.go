package followThreads

import (
	"context"
	"infion-BE/businesses/followThreads"

	"gorm.io/gorm"
)

type mysqlFollowThreadsRepository struct {
	Conn *gorm.DB
}

func NewFollowThreadsRepository(conn *gorm.DB) followThreads.Repository {
	return &mysqlFollowThreadsRepository{
		Conn: conn,
	}
}

func (nr *mysqlFollowThreadsRepository) Store(ctx context.Context, followThreadsDomain *followThreads.Domain) (followThreads.Domain, error) {
	rec := fromDomain(followThreadsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return followThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowThreadsRepository) GetByID(ctx context.Context, followThreadsId int) (followThreads.Domain, error) {
	rec := FollowThreads{}
	err := nr.Conn.Where("id = ?", followThreadsId).First(&rec).Error
	if err != nil {
		return followThreads.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlFollowThreadsRepository) Update(ctx context.Context, followThreadsDomain *followThreads.Domain) (followThreads.Domain, error) {
	rec := fromDomain(followThreadsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return followThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowThreadsRepository) Delete(ctx context.Context, followThreadsDomain *followThreads.Domain) (followThreads.Domain, error) {
	rec := fromDomain(followThreadsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return followThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFollowThreadsRepository) CountByThreadID(ctx context.Context,id int) (int, error) {
	rec := FollowThreads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("thread_id = ?", id).Where("status = ?", true).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlFollowThreadsRepository) CountByUserID(ctx context.Context,id uint) (int, error) {
	rec := FollowThreads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("user_id = ?", id).Where("status = ?", true).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlFollowThreadsRepository) GetDuplicate(ctx context.Context, threadID int, userID int) (followThreads.Domain, error) {
	rec := FollowThreads{}
	err := nr.Conn.Where("thread_id = ? AND user_id = ?", threadID, userID).First(&rec).Error
	if err != nil {
		return followThreads.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlFollowThreadsRepository) GetFollowThreadsByThreadID(ctx context.Context, threadID int) ([]followThreads.Domain, error) {
	var recordFollowThread []FollowThreads
	
	result := nr.Conn.Unscoped().Where("thread_id = ?", threadID).Find(&recordFollowThread)
	if result.Error != nil {
		return []followThreads.Domain{}, result.Error
	}

	return ToDomainArray(recordFollowThread), nil
}