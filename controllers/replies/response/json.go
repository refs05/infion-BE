package response

import (
	"infion-BE/businesses/replies"
	"time"
)

type Replies struct {
	ID        int       `json:"id"`
	CommentID int       `json:"comment_id"`
	UserID    int       `json:"user_id"`
	User      string    `json:"username"`
	UrlImg    string    `json:"url_img"`
	Reply     string    `json:"reply"`
	LikeCount int       `json:"like_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain replies.Domain) Replies {
	return Replies{
		ID:        domain.ID,
		CommentID: domain.CommentID,
		UserID:    domain.UserID,
		User:      domain.User,
		UrlImg:    domain.UrlImg,
		Reply:     domain.Reply,
		LikeCount: domain.LikeCount,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func NewResponseArray(domainRelpies []replies.Domain) []Replies {
	var resp []Replies

	for _, value := range domainRelpies {
		resp = append(resp, FromDomain(value))
	}

	return resp
}
