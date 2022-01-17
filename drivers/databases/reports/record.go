package reports

import (
	reportsUsecase "infion-BE/businesses/reports"
	"infion-BE/drivers/databases/threads"
	"infion-BE/drivers/databases/users"
	"time"
)

type Reports struct {
	ID           	int `gorm:"primaryKey"`
	ThreadID		int
	Thread			threads.Threads
	UserID			int
	User			users.User
	ReportMessage	string
	Status			string
	CreatedAt    	time.Time `gorm:"<-:create"`
	UpdatedAt   	time.Time
}

func fromDomain(domain *reportsUsecase.Domain) *Reports {
	return &Reports{
		ID: 			domain.ID,
		ThreadID: 		domain.ThreadID,
		UserID: 		domain.UserID,
		ReportMessage: 	domain.ReportMessage,
		Status: 		domain.Status,
	}
}

func (rec *Reports) toDomain() reportsUsecase.Domain {
	return reportsUsecase.Domain{
		ID:           	rec.ID,
		ThreadID: 		rec.ThreadID,
		Title: 			rec.Thread.Title,
		UserID: 		rec.UserID,
		Moderator: 		rec.User.Username,
		ReportMessage: 	rec.ReportMessage,
		Status: 		rec.Status,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}

func ToDomainArray(modelReports []Reports) []reportsUsecase.Domain {
	var response []reportsUsecase.Domain

	for _, val := range modelReports{
		response = append(response, val.toDomain())
	}
	return response
}