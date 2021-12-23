package response

import (
	"infion-BE/businesses/threads"
	"time"
)

type Threads struct {
	ID           	int `json:"id"`
	Title      		string `json:"title"`
	Img				string `json:"img"`
	Content      	string `json:"content"`
	Category   		string `json:"category"`
	UserID			int `json:"user_id"`
	LikeCount		int `json:"like_count"`
	CommentCount	int `json:"comment_count"`
	CreatedAt    	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func FromDomain(domain threads.Domain) Threads {
	return Threads{
		ID:           	domain.ID,
		Title:      	domain.Title,
		Img: 			domain.Img,
		Content:    	domain.Content,
		Category:   	domain.Category,
		UserID: 		domain.UserID,
		LikeCount:  	domain.LikeCount,
		CommentCount: 	domain.CommentCount,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:    	domain.UpdatedAt,
	}
}

func NewResponseArray(domainThreads []threads.Domain) []Threads {
	var resp []Threads

	for _, value := range domainThreads {
		resp = append(resp, FromDomain(value))
	}

	return resp
}