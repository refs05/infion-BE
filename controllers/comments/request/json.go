package request

import "infion-BE/businesses/comments"

type Comments struct {
	ThreadID		int `json:"thread_id"`
	UserID			int `json:"user_id"`
	Comment			string `json:"comment"`
	LikeCount		int `json:"like_count"`
	ReplyCount		int `json:"reply_count"`
}

func (req *Comments) ToDomain() *comments.Domain {
	return &comments.Domain{
		ThreadID: 		req.ThreadID,
		UserID: 		req.UserID,
		Comment: 		req.Comment,
		LikeCount:  	req.LikeCount,
		ReplyCount: 	req.ReplyCount,
	}
}