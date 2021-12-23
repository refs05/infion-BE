package threads

import (
	threadsUsecase "infion-BE/businesses/threads"
	"infion-BE/drivers/databases/users"
	"time"
)

type Threads struct {
	ID           	int `gorm:"primaryKey"`
	Title      		string 
	Img				string
	Content      	string
	Category   		string
	UserID			int
	User			users.User
	LikeCount		int
	CommentCount	int
	CreatedAt    	time.Time
	UpdatedAt   	time.Time
}

func fromDomain(domain *threadsUsecase.Domain) *Threads {
	return &Threads{
		ID:           	domain.ID,
		Title:      	domain.Title,
		Img: 			domain.Img,
		Content:    	domain.Content,
		Category:   	domain.Category,
		UserID: 		domain.UserID,
		LikeCount:  	domain.LikeCount,
		CommentCount: 	domain.CommentCount,
	}
}

func (rec *Threads) toDomain() threadsUsecase.Domain {
	return threadsUsecase.Domain{
		ID:           	rec.ID,
		Title:      	rec.Title,
		Img: 			rec.Img,
		Content:    	rec.Content,
		Category:   	rec.Category,
		UserID: 		rec.UserID,
		User:			rec.User.Username,
		LikeCount:  	rec.LikeCount,
		CommentCount: 	rec.CommentCount,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}

func ToDomain(rec Threads) threadsUsecase.Domain {
	return threadsUsecase.Domain{
		ID:           	rec.ID,
		Title:      	rec.Title,
		Img: 			rec.Img,
		Content:    	rec.Content,
		Category:   	rec.Category,
		UserID: 		rec.UserID,
		LikeCount:  	rec.LikeCount,
		CommentCount: 	rec.CommentCount,
		CreatedAt:    	rec.CreatedAt,
		UpdatedAt:    	rec.UpdatedAt,
	}
}

func ToDomainArray(modelThreads []Threads) []threadsUsecase.Domain {
	var response []threadsUsecase.Domain

	for _, val := range modelThreads{
		response = append(response, ToDomain(val))
	}
	return response
}