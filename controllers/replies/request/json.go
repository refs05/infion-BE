package request

import "infion-BE/businesses/replies"

type Replies struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	Reply     string `json:"reply"`
	LikeCount int    `json:"like_count"`
}

func (req *Replies) ToDomain() *replies.Domain {
	return &replies.Domain{
		CommentID: req.CommentID,
		UserID:    req.UserID,
		Reply:     req.Reply,
		LikeCount: req.LikeCount,
	}
}
