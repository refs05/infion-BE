package request

import "infion-BE/businesses/followThreads"

type FollowThreads struct {
	ID          	int `json:"id"`
	ThreadID		int	`json:"thread_id"`
	UserID			int	`json:"user_id"`
	Status			bool `json:"status"`
}

func (req *FollowThreads) ToDomain() *followThreads.Domain {
	return &followThreads.Domain{
		ID:				req.ID,
		ThreadID:		req.ThreadID,
		UserID: 		req.UserID,
		Status: 		req.Status,
	}
}