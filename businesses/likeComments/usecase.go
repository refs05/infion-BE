package likeComments

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type likeCommentsUsecase struct {
	likeCommentsRepository  Repository
	contextTimeout  time.Duration
}

func NewLikeCommentsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &likeCommentsUsecase{
		likeCommentsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *likeCommentsUsecase) Store(ctx context.Context, likeCommentsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	duplicate, err := tu.likeCommentsRepository.GetDuplicate(ctx, likeCommentsDomain.CommentID, likeCommentsDomain.UserID)
	if err == nil {
		duplicate.Status = !duplicate.Status
		tu.likeCommentsRepository.Update(ctx, &duplicate)
		return duplicate, nil
	}

	result, err := tu.likeCommentsRepository.Store(ctx, likeCommentsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *likeCommentsUsecase) GetByID(ctx context.Context, likeCommentsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if likeCommentsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.likeCommentsRepository.GetByID(ctx, likeCommentsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *likeCommentsUsecase) Update(ctx context.Context, likeCommentsDomain *Domain) (*Domain, error) {
	existedLikeComments, err := tu.likeCommentsRepository.GetByID(ctx, likeCommentsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeCommentsDomain.ID = existedLikeComments.ID

	result, err := tu.likeCommentsRepository.Update(ctx, likeCommentsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *likeCommentsUsecase) Delete(ctx context.Context, likeCommentsDomain *Domain) (*Domain, error) {
	existedLikeComments, err := tu.likeCommentsRepository.GetByID(ctx, likeCommentsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeCommentsDomain.ID = existedLikeComments.ID

	result, err := tu.likeCommentsRepository.Delete(ctx, likeCommentsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}