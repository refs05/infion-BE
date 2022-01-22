package likeComments

import (
	likeCommentsUsecase "infion-BE/businesses/likeComments"
	"time"
)

type LikeComments struct {
	ID        int `gorm:"primaryKey"`
	CommentID int
	UserID    int
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func fromDomain(domain *likeCommentsUsecase.Domain) *LikeComments {
	return &LikeComments{
		ID:        domain.ID,
		CommentID: domain.CommentID,
		UserID:    domain.UserID,
	}
}

func (rec *LikeComments) toDomain() likeCommentsUsecase.Domain {
	return likeCommentsUsecase.Domain{
		ID:        rec.ID,
		CommentID: rec.CommentID,
		UserID:    rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainArray(modelLikeComments []LikeComments) []likeCommentsUsecase.Domain {
	var response []likeCommentsUsecase.Domain

	for _, val := range modelLikeComments{
		response = append(response, val.toDomain())
	}
	return response
}
