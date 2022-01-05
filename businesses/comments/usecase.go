package comments

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/replies"
	"time"
)

type commentsUsecase struct {
	commentsRepository  Repository
	repliesRepository	replies.Repository
	contextTimeout  time.Duration
}

func NewCommentsUsecase(tr Repository, timeout time.Duration, rr replies.Repository) Usecase {
	return &commentsUsecase{
		commentsRepository:  tr,
		contextTimeout:  timeout,
		repliesRepository: rr,
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

	res.Replies, _ = tu.repliesRepository.GetRepliesByCommentID(ctx, commentsId)

	return res, nil
}

func (tu *commentsUsecase) GetComments(ctx context.Context) ([]Domain, error) {
	result, err := tu.commentsRepository.GetComments(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (tu *commentsUsecase) GetCommentsByThreadID(ctx context.Context, threadId int) ([]Domain, error) {
	result, err := tu.commentsRepository.GetCommentsByThreadID(ctx, threadId)
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