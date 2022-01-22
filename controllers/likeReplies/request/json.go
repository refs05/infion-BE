package request

import "infion-BE/businesses/likeReplies"

type LikeReplies struct {
	ID          	int `json:"id"`
	ReplyID			int	`json:"reply_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
}

func (req *LikeReplies) ToDomain() *likeReplies.Domain {
	return &likeReplies.Domain{
		ID:				req.ID,
		ReplyID:		req.ReplyID,
		UserID: 		req.UserID,
		Status: 		req.Status,
	}
}