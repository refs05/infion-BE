package likeThreads

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type likeThreadsUsecase struct {
	likeThreadsRepository  Repository
	contextTimeout  time.Duration
}

func NewLikeThreadsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &likeThreadsUsecase{
		likeThreadsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *likeThreadsUsecase) Store(ctx context.Context, likeThreadsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	duplicate, err := tu.likeThreadsRepository.GetDuplicate(ctx, likeThreadsDomain.ThreadID, likeThreadsDomain.UserID)
	if err == nil {
		duplicate.Status = !duplicate.Status
		tu.likeThreadsRepository.Update(ctx, &duplicate)
		return duplicate, nil
	}

	result, err := tu.likeThreadsRepository.Store(ctx, likeThreadsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *likeThreadsUsecase) GetByID(ctx context.Context, likeThreadsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if likeThreadsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.likeThreadsRepository.GetByID(ctx, likeThreadsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *likeThreadsUsecase) Update(ctx context.Context, likeThreadsDomain *Domain) (*Domain, error) {
	existedLikeThreads, err := tu.likeThreadsRepository.GetByID(ctx, likeThreadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeThreadsDomain.ID = existedLikeThreads.ID

	result, err := tu.likeThreadsRepository.Update(ctx, likeThreadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *likeThreadsUsecase) Delete(ctx context.Context, likeThreadsDomain *Domain) (*Domain, error) {
	existedLikeThreads, err := tu.likeThreadsRepository.GetByID(ctx, likeThreadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	likeThreadsDomain.ID = existedLikeThreads.ID

	result, err := tu.likeThreadsRepository.Delete(ctx, likeThreadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}