package request

import "infion-BE/businesses/likeThreads"

type LikeThreads struct {
	ID          	int `json:"id"`
	ThreadID		int	`json:"thread_id"`
	UserID			int	`json:"user_id"`
}

func (req *LikeThreads) ToDomain() *likeThreads.Domain {
	return &likeThreads.Domain{
		ID:				req.ID,
		ThreadID:		req.ThreadID,
		UserID: 		req.UserID,
	}
}