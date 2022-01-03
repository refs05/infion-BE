package comments

import (
	"context"
	"infion-BE/businesses/comments"

	"gorm.io/gorm"
)

type mysqlCommentsRepository struct {
	Conn *gorm.DB
}

func NewCommentsRepository(conn *gorm.DB) comments.Repository {
	return &mysqlCommentsRepository{
		Conn: conn,
	}
}

func (nr *mysqlCommentsRepository) Store(ctx context.Context, commentsDomain *comments.Domain) (comments.Domain, error) {
	rec := fromDomain(commentsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return comments.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlCommentsRepository) GetByID(ctx context.Context, commentsId int) (comments.Domain, error) {
	rec := Comments{}
	err := nr.Conn.Where("comments.id = ?", commentsId).Joins("Thread").Joins("User").First(&rec).Error
	if err != nil {
		return comments.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlCommentsRepository) Update(ctx context.Context, commentsDomain *comments.Domain) (comments.Domain, error) {
	rec := fromDomain(commentsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return comments.Domain{}, err
	}

	return record, nil
}

func (nr *mysqlCommentsRepository) Delete(ctx context.Context, commentsDomain *comments.Domain) (comments.Domain, error) {
	rec := fromDomain(commentsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlCommentsRepository) GetComments(ctx context.Context) ([]comments.Domain, error) {
	var recordThread []Comments
	
	result := nr.Conn.Unscoped().Joins("Thread").Joins("User").Find(&recordThread)
	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToDomainArray(recordThread), nil
}