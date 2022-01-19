package followUsers

import (
	followUsersUsecase "infion-BE/businesses/followUsers"
	"time"
)

type FollowUsers struct {
	ID        	int `gorm:"primaryKey"`
	FollowedID	int
	FollowerID	int
	CreatedAt 	time.Time `gorm:"<-:create"`
	UpdatedAt 	time.Time
}

func fromDomain(domain *followUsersUsecase.Domain) *FollowUsers {
	return &FollowUsers{
		ID:			domain.ID,
		FollowedID: domain.FollowedID,
		FollowerID:	domain.FollowerID,
	}
}

func (rec *FollowUsers) toDomain() followUsersUsecase.Domain {
	return followUsersUsecase.Domain{
		ID:				rec.ID,
		FollowedID:		rec.FollowedID,
		FollowerID:		rec.FollowerID,
		CreatedAt:		rec.CreatedAt,
		UpdatedAt:		rec.UpdatedAt,
	}
}
