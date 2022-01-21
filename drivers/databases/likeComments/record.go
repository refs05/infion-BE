package likeComments

import (
	likeCommentsUsecase "infion-BE/businesses/likeComments"
	"time"
)

type LikeComments struct {
	ID        int `gorm:"primaryKey"`
	CommentID int
	UserID    int
	Status	  bool `gorm:"default:true"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func fromDomain(domain *likeCommentsUsecase.Domain) *LikeComments {
	return &LikeComments{
		ID:        domain.ID,
		CommentID: domain.CommentID,
		UserID:    domain.UserID,
		Status:    domain.Status,
	}
}

func (rec *LikeComments) toDomain() likeCommentsUsecase.Domain {
	return likeCommentsUsecase.Domain{
		ID:        rec.ID,
		CommentID: rec.CommentID,
		UserID:    rec.UserID,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
