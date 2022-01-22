package replies

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeReplies"
	"time"
)

type repliesUsecase struct {
	repliesRepository		Repository
	contextTimeout			time.Duration
	likeRepliesRepository	likeReplies.Repository
}

func NewRepliesUsecase(tr Repository, timeout time.Duration, lrr likeReplies.Repository) Usecase {
	return &repliesUsecase{
		repliesRepository: tr,
		contextTimeout:    timeout,
		likeRepliesRepository: lrr,
	}
}

func (tu *repliesUsecase) Store(ctx context.Context, repliesDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.repliesRepository.Store(ctx, repliesDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *repliesUsecase) GetByID(ctx context.Context, repliesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if repliesId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.repliesRepository.GetByID(ctx, repliesId)
	if err != nil {
		return Domain{}, err
	}

	res.LikeCount, err = tu.likeRepliesRepository.CountByReplyID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}

	_, err = tu.repliesRepository.Update(ctx, &res)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *repliesUsecase) GetReplies(ctx context.Context) ([]Domain, error) {
	result, err := tu.repliesRepository.GetReplies(ctx)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeRepliesRepository.CountByReplyID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.repliesRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *repliesUsecase) GetRepliesByCommentID(ctx context.Context, threadId int) ([]Domain, error) {
	result, err := tu.repliesRepository.GetRepliesByCommentID(ctx, threadId)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeRepliesRepository.CountByReplyID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.repliesRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}
	
	return result, nil
}

func (tu *repliesUsecase) Update(ctx context.Context, repliesDomain *Domain) (*Domain, error) {
	existedReplies, err := tu.repliesRepository.GetByID(ctx, repliesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	repliesDomain.ID = existedReplies.ID

	result, err := tu.repliesRepository.Update(ctx, repliesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *repliesUsecase) Delete(ctx context.Context, repliesDomain *Domain) (*Domain, error) {
	existedReplies, err := tu.repliesRepository.GetByID(ctx, repliesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	repliesDomain.ID = existedReplies.ID

	likeReplies, _ := tu.likeRepliesRepository.GetLikeRepliesByReplyID(ctx, repliesDomain.ID)
	for i := range likeReplies {
		_, _ = tu.likeRepliesRepository.Delete(ctx, &likeReplies[i])
	}


	result, err := tu.repliesRepository.Delete(ctx, repliesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
