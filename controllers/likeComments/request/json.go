package request

import "infion-BE/businesses/likeComments"

type LikeComments struct {
	ID          	int `json:"id"`
	CommentID		int	`json:"comment_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
}

func (req *LikeComments) ToDomain() *likeComments.Domain {
	return &likeComments.Domain{
		ID:				req.ID,
		CommentID:		req.CommentID,
		UserID: 		req.UserID,
		Status: 		req.Status,
	}
}