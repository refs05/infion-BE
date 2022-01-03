package comments

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type commentsUsecase struct {
	commentsRepository  Repository
	contextTimeout  time.Duration
}

func NewCommentsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &commentsUsecase{
		commentsRepository:  tr,
		contextTimeout:  timeout,
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

	return res, nil
}

func (tu *commentsUsecase) GetComments(ctx context.Context) ([]Domain, error) {
	result, err := tu.commentsRepository.GetComments(ctx)
	if err != nil {
		return []Domain{}, err
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

	result, err := tu.commentsRepository.Delete(ctx, commentsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}