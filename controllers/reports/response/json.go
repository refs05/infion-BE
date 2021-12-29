package response

import (
	"infion-BE/businesses/reports"
	"time"
)

type Reports struct {
	ID           	int `json:"id"`
	ThreadID		int	`json:"thread_id"`
	Title			string `json:"title"`
	UserID			int `json:"user_id"`
	Moderator		string `json:"moderator"`
	ReportMessage   string `json:"report_message"`
	Status   		string `json:"status"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain reports.Domain) Reports {
	return Reports{
		ID: 			domain.ID,
		ThreadID: 		domain.ThreadID,
		Title: 			domain.Title,
		UserID: 		domain.UserID,
		Moderator: 		domain.Moderator,
		ReportMessage: 	domain.ReportMessage,
		Status: 		domain.Status,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}

func NewResponseArray(domainReports []reports.Domain) []Reports {
	var resp []Reports

	for _, value := range domainReports {
		resp = append(resp, FromDomain(value))
	}

	return resp
}