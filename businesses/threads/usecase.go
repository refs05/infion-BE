package threads

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/comments"
	"infion-BE/businesses/likeThreads"

	"time"
)

type threadsUsecase struct {
	threadsRepository		Repository
	contextTimeout			time.Duration
	likeThreadsRepository	likeThreads.Repository
	commentsRepository		comments.Repository
}

func NewThreadsUsecase(tr Repository, timeout time.Duration, ltr likeThreads.Repository, cr comments.Repository) Usecase {
	return &threadsUsecase{
		threadsRepository:  tr,
		contextTimeout:  timeout,
		likeThreadsRepository: ltr,
		commentsRepository: cr,
	}
}

func (tu *threadsUsecase) Store(ctx context.Context, threadsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.threadsRepository.Store(ctx, threadsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *threadsUsecase) GetByID(ctx context.Context, threadsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if threadsId <= 0 {
		return Domain{}, businesses.ErrThreadsIDResource
	}
	res, err := tu.threadsRepository.GetByID(ctx, threadsId)
	if err != nil {
		return Domain{}, err
	}

	res.LikeCount, err = tu.likeThreadsRepository.CountByThreadID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}
	
	res.CommentCount, err = tu.commentsRepository.CountByThreadID(ctx, res.ID)
	if err != nil {
		return Domain{}, err
	}

	_, err = tu.threadsRepository.Update(ctx, &res)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *threadsUsecase) GetThreads(ctx context.Context) ([]Domain, error) {
	result, err := tu.threadsRepository.GetThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeThreadsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].CommentCount, err = tu.commentsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.threadsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *threadsUsecase) GetThreadsByCategory(ctx context.Context, category string) ([]Domain, error) {
	result, err := tu.threadsRepository.GetThreadsByCategory(ctx, category)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeThreadsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].CommentCount, err = tu.commentsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.threadsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *threadsUsecase) GetThreadsBySort(ctx context.Context, sort string) ([]Domain, error) {
	result, err := tu.threadsRepository.GetThreadsBySort(ctx, sort)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeThreadsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].CommentCount, err = tu.commentsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.threadsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}

	return result, nil
}

func (tu *threadsUsecase) GetThreadsBySortCategory(ctx context.Context, sort string, category string) ([]Domain, error) {
	result, err := tu.threadsRepository.GetThreadsBySortCategory(ctx, sort, category)
	if err != nil {
		return []Domain{}, err
	}

	for i := range result {
		result[i].LikeCount, err = tu.likeThreadsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		result[i].CommentCount, err = tu.commentsRepository.CountByThreadID(ctx, result[i].ID)
		if err != nil {
			return []Domain{}, err
		}
	}

	for i := range result {
		_, err = tu.threadsRepository.Update(ctx, &result[i])
		if err != nil {
			return []Domain{}, err
		}
	}
	
	return result, nil
}

func (tu *threadsUsecase) Update(ctx context.Context, threadsDomain *Domain) (*Domain, error) {
	existedThreads, err := tu.threadsRepository.GetByID(ctx, threadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	threadsDomain.ID = existedThreads.ID
	threadsDomain.LikeCount = existedThreads.LikeCount

	result, err := tu.threadsRepository.Update(ctx, threadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *threadsUsecase) Delete(ctx context.Context, threadsDomain *Domain) (*Domain, error) {
	existedThreads, err := tu.threadsRepository.GetByID(ctx, threadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	threadsDomain.ID = existedThreads.ID

	result, err := tu.threadsRepository.Delete(ctx, threadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}