package comments

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeComments"
	"infion-BE/businesses/replies"
	"time"
)

type commentsUsecase struct {
	commentsRepository  	Repository
	contextTimeout  		time.Duration
	repliesRepository 		replies.Repository
	likeCommentsRepository	likeComments.Repository
}

func NewCommentsUsecase(tr Repository, timeout time.Duration, rr replies.Repository, lcr likeComments.Repository) Usecase {
	return &commentsUsecase{
		commentsRepository:  tr,
		contextTimeout:  timeout,
		repliesRepository: rr,
		likeCommentsRepository: lcr,
	}
}

func (tu *commentsUsecase) Store(ctx context.Context, commentsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.commentsRepository.Store(ctx, commentsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *commentsUsecase) GetByID(ctx context.Context, commentsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if commentsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.commentsRepository.GetByID(ctx, commentsId)
	if err != nil {
		return Domain{}, err
	}

	res.Replies, err = tu.repliesRepository.GetRepliesByCommentID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}

	res.LikeCount, err = tu.likeCommentsRepository.CountByCommentID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}

	res.ReplyCount, err = tu.repliesRepository.CountByCommentID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}

	_, err = tu.commentsRepository.Update(ctx, &res)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *commentsUsecase) GetComments(ctx context.Context) ([]Domain, error) {
	result, err := tu.commentsRepository.GetComments(ctx)
	if err != nil {
		return []Domain{}, err
	}
	for i := range result {
		result[i].Replies, err = tu.repliesRepository.GetRepliesByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeCommentsRepository.CountByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].ReplyCount, err =  tu.repliesRepository.CountByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.commentsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *commentsUsecase) GetCommentsByThreadID(ctx context.Context, threadId int) ([]Domain, error) {
	result, err := tu.commentsRepository.GetCommentsByThreadID(ctx, threadId)
	if err != nil {
		return []Domain{}, err
	}
	for i := range result {
		result[i].Replies, err = tu.repliesRepository.GetRepliesByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeCommentsRepository.CountByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].ReplyCount, err =  tu.repliesRepository.CountByCommentID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.commentsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *commentsUsecase) Update(ctx context.Context, commentsDomain *Domain) (*Domain, error) {
	existedComments, err := tu.commentsRepository.GetByID(ctx, commentsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	commentsDomain.ID = existedComments.ID

	result, err := tu.commentsRepository.Update(ctx, commentsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *commentsUsecase) Delete(ctx context.Context, commentsDomain *Domain) (*Domain, error) {
	existedComments, err := tu.commentsRepository.GetByID(ctx, commentsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	commentsDomain.ID = existedComments.ID

	replies, _ := tu.repliesRepository.GetRepliesByCommentID(ctx, commentsDomain.ID)
	for i := range replies {
		_, err = tu.repliesRepository.Delete(ctx, &replies[i])
		if err != nil {
			return &Domain{}, err
		}
	}

	result, err := tu.commentsRepository.Delete(ctx, commentsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}