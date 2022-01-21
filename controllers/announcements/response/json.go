package response

import (
	"infion-BE/businesses/announcements"
	"time"
)

type Announcements struct {
	ID           	int `json:"id"`
	UserID			int `json:"announcer_id"`
	Announcer		string `json:"announcer"`
	Message   		string `json:"message"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain announcements.Domain) Announcements {
	return Announcements{
		ID: 			domain.ID,
		UserID: 		domain.UserID,
		Announcer: 		domain.Announcer,
		Message: 		domain.Message,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}

func NewResponseArray(domainAnnouncements []announcements.Domain) []Announcements {
	var resp []Announcements

	for _, value := range domainAnnouncements {
		resp = append(resp, FromDomain(value))
	}

	return resp
}