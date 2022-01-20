package threads

import (
	"context"
	"infion-BE/businesses/threads"

	"gorm.io/gorm"
)

type mysqlThreadsRepository struct {
	Conn *gorm.DB
}

func NewThreadsRepository(conn *gorm.DB) threads.Repository {
	return &mysqlThreadsRepository{
		Conn: conn,
	}
}

func (nr *mysqlThreadsRepository) Store(ctx context.Context, threadsDomain *threads.Domain) (threads.Domain, error) {
	rec := fromDomain(threadsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return threads.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlThreadsRepository) GetByID(ctx context.Context, threadsId int) (threads.Domain, error) {
	rec := Threads{}
	err := nr.Conn.Where("threads.id = ?", threadsId).Joins("User").First(&rec).Error
	if err != nil {
		return threads.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlThreadsRepository) Update(ctx context.Context, threadsDomain *threads.Domain) (threads.Domain, error) {
	rec := fromDomain(threadsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return threads.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlThreadsRepository) Delete(ctx context.Context, threadsDomain *threads.Domain) (threads.Domain, error) {
	rec := fromDomain(threadsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlThreadsRepository) GetThreads(ctx context.Context) ([]threads.Domain, error) {
	var recordThread []Threads
	
	result := nr.Conn.Unscoped().Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlThreadsRepository) GetThreadsByCategory(ctx context.Context, category string) ([]threads.Domain, error) {
	var recordThread []Threads
	
	result := nr.Conn.Unscoped().Where("threads.category = ?", category).Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlThreadsRepository) GetThreadsBySort(ctx context.Context, sort string) ([]threads.Domain, error) {
	var recordThread []Threads
	
	result := nr.Conn.Unscoped().Order(sort).Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlThreadsRepository) GetThreadsBySortCategory(ctx context.Context, sort string, category string) ([]threads.Domain, error) {
	var recordThread []Threads
	
	result := nr.Conn.Unscoped().Order(sort).Where("threads.category = ?", category).Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlThreadsRepository) GetThreadsByUserID(ctx context.Context, userID int) ([]threads.Domain, error) {
	var recordThread []Threads
	
	result := nr.Conn.Unscoped().Where("threads.user_id = ?", userID).Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}

func (nr *mysqlThreadsRepository) CountByUserID(ctx context.Context, id uint) (int, error) {
	rec := Threads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("user_id = ?", id).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlThreadsRepository) GetThreadLikeCountByUserID(ctx context.Context, userID uint) (int, error) {
	rec := Threads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("threads.user_id = ?", userID).Joins("LikeThreads").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlThreadsRepository) GetThreadFollowerCountByUserID(ctx context.Context, userID uint) (int, error) {
	rec := Threads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("threads.user_id = ?", userID).Joins("FollowThreads").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}