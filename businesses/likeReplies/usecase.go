package likeReplies

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type likeRepliesUsecase struct {
	likeRepliesRepository  Repository
	contextTimeout  time.Duration
}

func NewLikeRepliesUsecase(tr Repository, timeout time.Duration) Usecase {
	return &likeRepliesUsecase{
		likeRepliesRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *likeRepliesUsecase) Store(ctx context.Context, likeRepliesDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	duplicate, err := tu.likeRepliesRepository.GetDuplicate(ctx, likeRepliesDomain.ReplyID, likeRepliesDomain.UserID)
	if err == nil {
		duplicate.Status = !duplicate.Status
		tu.likeRepliesRepository.Update(ctx, &duplicate)
		return duplicate, nil
	}

	result, err := tu.likeRepliesRepository.Store(ctx, likeRepliesDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *likeRepliesUsecase) GetByID(ctx context.Context, likeRepliesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if likeRepliesId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.likeRepliesRepository.GetByID(ctx, likeRepliesId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *likeRepliesUsecase) Update(ctx context.Context, likeRepliesDomain *Domain) (*Domain, error) {
	existedLikeReplies, err := tu.likeRepliesRepository.GetByID(ctx, likeRepliesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeRepliesDomain.ID = existedLikeReplies.ID

	result, err := tu.likeRepliesRepository.Update(ctx, likeRepliesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *likeRepliesUsecase) Delete(ctx context.Context, likeRepliesDomain *Domain) (*Domain, error) {
	existedLikeReplies, err := tu.likeRepliesRepository.GetByID(ctx, likeRepliesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeRepliesDomain.ID = existedLikeReplies.ID

	result, err := tu.likeRepliesRepository.Delete(ctx, likeRepliesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}