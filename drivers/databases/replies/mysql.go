package replies

import (
	"context"
	"infion-BE/businesses/replies"

	"gorm.io/gorm"
)

type mysqlRepliesRepository struct {
	Conn *gorm.DB
}

func NewRepliesRepository(conn *gorm.DB) replies.Repository {
	return &mysqlRepliesRepository{
		Conn: conn,
	}
}

func (nr *mysqlRepliesRepository) Store(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	rec := fromDomain(repliesDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return replies.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return replies.Domain{}, err
	}

	return record, nil
}

func (nr mysqlRepliesRepository) GetByID(ctx context.Context, repliesId int) (replies.Domain, error) {
	rec := Replies{}
	err := nr.Conn.Where("replies.id = ?", repliesId).Joins("User").First(&rec).Error
	if err != nil {
		return replies.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr mysqlRepliesRepository) Update(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	rec := fromDomain(repliesDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return replies.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return replies.Domain{}, err
	}

	return record, nil
}

func (nr mysqlRepliesRepository) Delete(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	rec := fromDomain(repliesDomain)

	result := nr.Conn.Delete(&rec)
	if result.Error != nil {
		return replies.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr mysqlRepliesRepository) GetReplies(ctx context.Context) ([]replies.Domain, error) {
	var recordReplies []Replies

	result := nr.Conn.Unscoped().Joins("User").Find(&recordReplies)
	if result.Error != nil {
		return []replies.Domain{}, result.Error
	}

	return ToDomainArray(recordReplies), nil
}

func (nr *mysqlRepliesRepository) GetRepliesByCommentID(ctx context.Context, commentId int) ([]replies.Domain, error) {
	var recordReplies []Replies
	
	result := nr.Conn.Where("replies.comment_id = ?", commentId).Joins("User").Find(&recordReplies)
	if result.Error != nil {
		return []replies.Domain{}, result.Error
	}

	return ToDomainArray(recordReplies), nil
}