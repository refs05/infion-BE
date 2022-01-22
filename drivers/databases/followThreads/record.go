package followThreads

import (
	followThreadsUsecase "infion-BE/businesses/followThreads"
	"time"
)

type FollowThreads struct {
	ID        int `gorm:"primaryKey"`
	ThreadID  int
	UserID    int
	Status	  bool `gorm:"default:true"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func fromDomain(domain *followThreadsUsecase.Domain) *FollowThreads {
	return &FollowThreads{
		ID:       domain.ID,
		ThreadID: domain.ThreadID,
		UserID:   domain.UserID,
		Status:   domain.Status,
	}
}

func (rec *FollowThreads) toDomain() followThreadsUsecase.Domain {
	return followThreadsUsecase.Domain{
		ID:        rec.ID,
		ThreadID:  rec.ThreadID,
		UserID:    rec.UserID,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainArray(modelFollowThreads []FollowThreads) []followThreadsUsecase.Domain {
	var response []followThreadsUsecase.Domain

	for _, val := range modelFollowThreads{
		response = append(response, val.toDomain())
	}
	return response
}