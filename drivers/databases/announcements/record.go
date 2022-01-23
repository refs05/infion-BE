package announcements

import (
	announcementsUsecase "infion-BE/businesses/announcements"
	"infion-BE/drivers/databases/users"
	"time"
)

type Announcements struct {
	ID           	int `gorm:"primaryKey"`
	UserID			int
	User			users.User
	Message			string
	Img				string
	CreatedAt    	time.Time `gorm:"<-:create"`
	UpdatedAt   	time.Time
}

func fromDomain(domain *announcementsUsecase.Domain) *Announcements {
	return &Announcements{
		ID: 			domain.ID,
		UserID: 		domain.UserID,
		Message: 		domain.Message,
		Img:			domain.Img,
	}
}

func (rec *Announcements) toDomain() announcementsUsecase.Domain {
	return announcementsUsecase.Domain{
		ID:           	rec.ID,
		UserID: 		rec.UserID,
		Announcer: 		rec.User.Username,
		Message: 		rec.Message,
		Img:			rec.Img,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}

func ToDomainArray(modelAnnouncements []Announcements) []announcementsUsecase.Domain {
	var response []announcementsUsecase.Domain

	for _, val := range modelAnnouncements{
		response = append(response, val.toDomain())
	}
	return response
}