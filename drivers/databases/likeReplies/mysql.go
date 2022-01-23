package likeReplies

import (
	"context"
	"infion-BE/businesses/likeReplies"

	"gorm.io/gorm"
)

type mysqlLikeRepliesRepository struct {
	Conn *gorm.DB
}

func NewLikeRepliesRepository(conn *gorm.DB) likeReplies.Repository {
	return &mysqlLikeRepliesRepository{
		Conn: conn,
	}
}

func (nr *mysqlLikeRepliesRepository) Store(ctx context.Context, likeRepliesDomain *likeReplies.Domain) (likeReplies.Domain, error) {
	rec := fromDomain(likeRepliesDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return likeReplies.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeRepliesRepository) GetByID(ctx context.Context, likeRepliesId int) (likeReplies.Domain, error) {
	rec := LikeReplies{}
	err := nr.Conn.Where("id = ?", likeRepliesId).First(&rec).Error
	if err != nil {
		return likeReplies.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeRepliesRepository) Update(ctx context.Context, likeRepliesDomain *likeReplies.Domain) (likeReplies.Domain, error) {
	rec := fromDomain(likeRepliesDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return likeReplies.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeRepliesRepository) Delete(ctx context.Context, likeRepliesDomain *likeReplies.Domain) (likeReplies.Domain, error) {
	rec := fromDomain(likeRepliesDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return likeReplies.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLikeRepliesRepository) CountByReplyID(ctx context.Context,id int) (int, error) {
	rec := LikeReplies{}
	var count int64
	
	result := nr.Conn.Model(&rec).Where("reply_id = ?", id).Where("status = ?", true).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (nr *mysqlLikeRepliesRepository) GetDuplicate(ctx context.Context, replyID int, userID int) (likeReplies.Domain, error) {
	rec := LikeReplies{}
	err := nr.Conn.Where("reply_id = ? AND user_id = ?", replyID, userID).First(&rec).Error
	if err != nil {
		return likeReplies.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLikeRepliesRepository) GetLikeRepliesByReplyID(ctx context.Context, replyID int) ([]likeReplies.Domain, error) {
	var recordLikeReply []LikeReplies
	
	result := nr.Conn.Unscoped().Where("like_replies.reply_id = ?", replyID).Joins("Reply").Joins("User").Find(&recordLikeReply)
	if result.Error != nil {
		return []likeReplies.Domain{}, result.Error
	}

	return ToDomainArray(recordLikeReply), nil
}