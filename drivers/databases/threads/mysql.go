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