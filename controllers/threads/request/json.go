package request

import "infion-BE/businesses/threads"

type Threads struct {
	Title      		string `json:"title"`
	Img				string `json:"img"`
	Content      	string `json:"content"`
	Category   		string `json:"category"`
	UserID			int `json:"userID"`
	LikeCount		int `json:"likeCount"`
	CommentCount	int `json:"commentCount"`
}

func (req *Threads) ToDomain() *threads.Domain {
	return &threads.Domain{
		Title:      	req.Title,
		Img: 			req.Img,
		Content:    	req.Content,
		Category:   	req.Category,
		UserID: 		req.UserID,
		LikeCount:  	req.LikeCount,
		CommentCount: 	req.CommentCount,
	}
}