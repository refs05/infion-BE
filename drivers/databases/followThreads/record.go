package followThreads

import (
	followThreadsUsecase "infion-BE/businesses/followThreads"
	"time"
)

type FollowThreads struct {
	ID        int `gorm:"primaryKey"`
	ThreadID  int
	UserID    int `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *followThreadsUsecase.Domain) *FollowThreads {
	return &FollowThreads{
		ID:       domain.ID,
		ThreadID: domain.ThreadID,
		UserID:   domain.UserID,
	}
}

func (rec *FollowThreads) toDomain() followThreadsUsecase.Domain {
	return followThreadsUsecase.Domain{
		ID:        rec.ID,
		ThreadID:  rec.ThreadID,
		UserID:    rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
