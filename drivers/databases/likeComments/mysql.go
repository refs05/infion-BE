package likeComments

import (
	"context"
	"infion-BE/businesses/likeComments"

	"gorm.io/gorm"
)

type mysqlLikeCommentsRepository struct {
	Conn *gorm.DB
}

func NewLikeCommentsRepository(conn *gorm.DB) likeComments.Repository {
	return &mysqlLikeCommentsRepository{
		Conn: conn,
	}
}

func (nr *mysqlLikeCommentsRepository) Store(ctx context.Context, likeCommentsDomain *likeComments.Domain) (likeComments.Domain, error) {
	rec := fromDomain(likeCommentsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return likeComments.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) GetByID(ctx context.Context, likeCommentsId int) (likeComments.Domain, error) {
	rec := LikeComments{}
	err := nr.Conn.Where("id = ?", likeCommentsId).First(&rec).Error
	if err != nil {
		return likeComments.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) Update(ctx context.Context, likeCommentsDomain *likeComments.Domain) (likeComments.Domain, error) {
	rec := fromDomain(likeCommentsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return likeComments.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) Delete(ctx context.Context, likeCommentsDomain *likeComments.Domain) (likeComments.Domain, error) {
	rec := fromDomain(likeCommentsDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return likeComments.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) CountByCommentID(ctx context.Context,id int) (int, error) {
	rec := LikeComments{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("comment_id = ?", id).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}