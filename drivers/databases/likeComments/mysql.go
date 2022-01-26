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

	result := nr.Conn.Unscoped().Delete(&rec)
	if result.Error != nil {
		return likeComments.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) CountByCommentID(ctx context.Context,id int) (int, error) {
	rec := LikeComments{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("comment_id = ?", id).Where("status = ?", true).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlLikeCommentsRepository) GetDuplicate(ctx context.Context, commentID int, userID int) (likeComments.Domain, error) {
	rec := LikeComments{}
	err := nr.Conn.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&rec).Error
	if err != nil {
		return likeComments.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeCommentsRepository) GetLikeCommentsByCommentID(ctx context.Context, commentID int) ([]likeComments.Domain, error) {
	var recordLikeComment []LikeComments
	
	result := nr.Conn.Unscoped().Where("like_comments.comment_id = ?", commentID).Joins("Comment").Joins("User").Find(&recordLikeComment)
	if result.Error != nil {
		return []likeComments.Domain{}, result.Error
	}

	return ToDomainArray(recordLikeComment), nil
}