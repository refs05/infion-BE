package likeThreads

import (
	likeThreadsUsecase "infion-BE/businesses/likeThreads"
	"time"
)

type LikeThreads struct {
	ID        int `gorm:"primaryKey"`
	ThreadID  int
	UserID    int
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func fromDomain(domain *likeThreadsUsecase.Domain) *LikeThreads {
	return &LikeThreads{
		ID:       domain.ID,
		ThreadID: domain.ThreadID,
		UserID:   domain.UserID,
	}
}

func (rec *LikeThreads) toDomain() likeThreadsUsecase.Domain {
	return likeThreadsUsecase.Domain{
		ID:        rec.ID,
		ThreadID:  rec.ThreadID,
		UserID:    rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainArray(modelLikeThreads []LikeThreads) []likeThreadsUsecase.Domain {
	var response []likeThreadsUsecase.Domain

	for _, val := range modelLikeThreads{
		response = append(response, val.toDomain())
	}
	return response
}