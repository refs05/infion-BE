package likeThreads

import (
	"context"
	"infion-BE/businesses/likeThreads"

	"gorm.io/gorm"
)

type mysqlLikeThreadsRepository struct {
	Conn *gorm.DB
}

func NewLikeThreadsRepository(conn *gorm.DB) likeThreads.Repository {
	return &mysqlLikeThreadsRepository{
		Conn: conn,
	}
}

func (nr *mysqlLikeThreadsRepository) Store(ctx context.Context, likeThreadsDomain *likeThreads.Domain) (likeThreads.Domain, error) {
	rec := fromDomain(likeThreadsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return likeThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeThreadsRepository) GetByID(ctx context.Context, likeThreadsId int) (likeThreads.Domain, error) {
	rec := LikeThreads{}
	err := nr.Conn.Where("id = ?", likeThreadsId).First(&rec).Error
	if err != nil {
		return likeThreads.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeThreadsRepository) Update(ctx context.Context, likeThreadsDomain *likeThreads.Domain) (likeThreads.Domain, error) {
	rec := fromDomain(likeThreadsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return likeThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeThreadsRepository) Delete(ctx context.Context, likeThreadsDomain *likeThreads.Domain) (likeThreads.Domain, error) {
	rec := fromDomain(likeThreadsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return likeThreads.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeThreadsRepository) CountByThreadID(ctx context.Context,id int) (int, error) {
	rec := LikeThreads{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("thread_id = ?", id).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlLikeThreadsRepository) GetDuplicate(ctx context.Context, threadID int, userID int) (likeThreads.Domain, error) {
	rec := LikeThreads{}
	err := nr.Conn.Where("thread_id = ? AND user_id = ?", threadID, userID).First(&rec).Error
	if err != nil {
		return likeThreads.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeThreadsRepository) GetLikeThreadsByThreadID(ctx context.Context, threadID int) ([]likeThreads.Domain, error) {
	var recordLikeThread []LikeThreads
	
	result := nr.Conn.Unscoped().Where("like_threads.thread_id = ?", threadID).Joins("Thread").Joins("User").Find(&recordLikeThread)
	if result.Error != nil {
		return []likeThreads.Domain{}, result.Error
	}

	return ToDomainArray(recordLikeThread), nil
}