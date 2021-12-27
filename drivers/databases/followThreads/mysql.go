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