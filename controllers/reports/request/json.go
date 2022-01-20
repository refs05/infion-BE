package request

import "infion-BE/businesses/reports"

type Reports struct {
	ThreadID		int	`json:"thread_id"`
	UserID			int `json:"user_id"`
	ReportMessage   string `json:"report_message"`
	Status   		bool `json:"status"`
}

func (req *Reports) ToDomain() *reports.Domain {
	return &reports.Domain{
		ThreadID: 		req.ThreadID,
		UserID: 		req.UserID,
		ReportMessage: 	req.ReportMessage,
		Status: 		req.Status,
	}
}