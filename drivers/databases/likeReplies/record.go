package likeReplies

import (
	likeRepliesUsecase "infion-BE/businesses/likeReplies"
	"time"
)

type LikeReplies struct {
	ID			int `gorm:"primaryKey"`
	ReplyID		int
	UserID		int `gorm:"unique"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func fromDomain(domain *likeRepliesUsecase.Domain) *LikeReplies {
	return &LikeReplies{
		ID:			domain.ID,
		ReplyID:	domain.ReplyID,
		UserID:		domain.UserID,
	}
}

func (rec *LikeReplies) toDomain() likeRepliesUsecase.Domain {
	return likeRepliesUsecase.Domain{
		ID:			rec.ID,
		ReplyID:	rec.ReplyID,
		UserID:		rec.UserID,
		CreatedAt:	rec.CreatedAt,
		UpdatedAt:	rec.UpdatedAt,
	}
}
