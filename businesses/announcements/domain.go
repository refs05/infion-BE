package announcements

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int `json:"id"`
	UserID			int `json:"announcer_id"`
	Announcer		string `json:"announcer"`
	Message			string `json:"message"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt    	time.Time `json:"updated_at"`
	DeletedAt		time.Time `json:"deleted_at"`
}

type Usecase interface {
	Store(ctx context.Context, announcementsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, announcementsId int) (Domain, error)
	GetAnnouncements(ctx context.Context) ([]Domain, error)
	GetAnnouncementsByUserID(ctx context.Context, userID int) ([]Domain, error)
	Update(ctx context.Context, announcementsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, announcementsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, announcementsDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, announcementsId int) (Domain, error)
	GetAnnouncements(ctx context.Context) ([]Domain, error)
	GetAnnouncementsByUserID(ctx context.Context, userID int) ([]Domain, error)
	Update(ctx context.Context, announcementsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, announcementsDomain *Domain) (Domain, error)
}