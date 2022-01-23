package request

import "infion-BE/businesses/announcements"

type Announcements struct {
	UserID			int `json:"announcer_id"`
	Message   		string `json:"message"`
	Img				string `json:"img"`
}

func (req *Announcements) ToDomain() *announcements.Domain {
	return &announcements.Domain{
		UserID: 		req.UserID,
		Message: 		req.Message,
		Img:			req.Img,
	}
}