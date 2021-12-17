package threads

import (
	"context"
	"infion-BE/businesses/threads"

	"gorm.io/gorm"
)

type mysqlThreadsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) threads.Repository {
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

	return rec.toDomain(), nil
}

func (nr *mysqlThreadsRepository) GetByID(ctx context.Context, threadsId int) (threads.Domain, error) {
	rec := Threads{}
	err := nr.Conn.Where("id = ?", threadsId).First(&rec).Error
	if err != nil {
		return threads.Domain{}, err
	}
	return rec.toDomain(), nil
}

// func (nr *mysqlThreadsRepository) GetByTitle(ctx context.Context, threadsTitle string) (threads.Domain, error) {
// 	rec := Threads{}
// 	err := nr.Conn.Where("title = ?", threadsTitle).First(&rec).Error
// 	if err != nil {
// 		return threads.Domain{}, err
// 	}
// 	return rec.toDomain(), nil
// }

// func (nr *mysqlThreadsRepository) Fetch(ctx context.Context, page, perpage int) ([]threads.Domain, int, error) {
// 	rec := []Threads{}

// 	offset := (page - 1) * perpage
// 	err := nr.Conn.Preload("categories").Offset(offset).Limit(perpage).Find(&rec).Error
// 	if err != nil {
// 		return []threads.Domain{}, 0, err
// 	}

// 	var totalData int64
// 	err = nr.Conn.Count(&totalData).Error
// 	if err != nil {
// 		return []threads.Domain{}, 0, err
// 	}

// 	var domainThreads []threads.Domain
// 	for _, value := range rec {
// 		domainThreads = append(domainThreads, value.toDomain())
// 	}
// 	return domainThreads, int(totalData), nil
// }

func (nr *mysqlThreadsRepository) Update(ctx context.Context, threadsDomain *threads.Domain) (threads.Domain, error) {
	rec := fromDomain(threadsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	// err := nr.Conn.Preload("Category").First(&rec, rec.ID).Error
	// if err != nil {
	// 	return threads.Domain{}, result.Error
	// }

	return rec.toDomain(), nil
}