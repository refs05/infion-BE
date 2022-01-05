package replies

import (
	repliesUsecase "infion-BE/businesses/replies"
	"infion-BE/drivers/databases/comments"
	"infion-BE/drivers/databases/users"
	"time"
)

type Replies struct {
	ID        int `gorm:"primaryKey"`
	CommentID int
	Comment   comments.Comments
	UserID    int
	User      users.User
	UrlImg    string
	Reply     string
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *repliesUsecase.Domain) *Replies {
	return &Replies{
		ID:        domain.ID,
		CommentID: domain.CommentID,
		UserID:    domain.UserID,
		Reply:     domain.Reply,
		LikeCount: domain.LikeCount,
	}
}

func (rec *Replies) toDomain() repliesUsecase.Domain {
	return repliesUsecase.Domain{
		ID:        rec.ID,
		CommentID: rec.CommentID,
		UserID:    rec.UserID,
		User:      rec.User.Username,
		UrlImg:    rec.User.UrlImg,
		Reply:     rec.Reply,
		LikeCount: rec.LikeCount,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainArray(modelReplies []Replies) []repliesUsecase.Domain {
	var response []repliesUsecase.Domain

	for _, val := range modelReplies {
		response = append(response, val.toDomain())
	}
	return response
}
