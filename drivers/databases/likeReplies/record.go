package likeReplies

import (
	likeRepliesUsecase "infion-BE/businesses/likeReplies"
	"time"
)

type LikeReplies struct {
	ID			int `gorm:"primaryKey"`
	ReplyID		int
	UserID		int
	CreatedAt	time.Time `gorm:"<-:create"`
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

func ToDomainArray(modelLikeReplies []LikeReplies) []likeRepliesUsecase.Domain {
	var response []likeRepliesUsecase.Domain

	for _, val := range modelLikeReplies{
		response = append(response, val.toDomain())
	}
	return response
}