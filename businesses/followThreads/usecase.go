package followThreads

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type followThreadsUsecase struct {
	followThreadsRepository  Repository
	contextTimeout  time.Duration
}

func NewFollowThreadsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &followThreadsUsecase{
		followThreadsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *followThreadsUsecase) Store(ctx context.Context, followThreadsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	duplicate, err := tu.followThreadsRepository.GetDuplicate(ctx, followThreadsDomain.ThreadID, followThreadsDomain.UserID)
	if err == nil {
		duplicate.Status = !duplicate.Status
		tu.followThreadsRepository.Update(ctx, &duplicate)
		return duplicate, nil
	}

	result, err := tu.followThreadsRepository.Store(ctx, followThreadsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *followThreadsUsecase) GetByID(ctx context.Context, followThreadsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if followThreadsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.followThreadsRepository.GetByID(ctx, followThreadsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *followThreadsUsecase) Update(ctx context.Context, followThreadsDomain *Domain) (*Domain, error) {
	existedFollowThreads, err := tu.followThreadsRepository.GetByID(ctx, followThreadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	followThreadsDomain.ID = existedFollowThreads.ID

	result, err := tu.followThreadsRepository.Update(ctx, followThreadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *followThreadsUsecase) Delete(ctx context.Context, followThreadsDomain *Domain) (*Domain, error) {
	existedFollowThreads, err := tu.followThreadsRepository.GetByID(ctx, followThreadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	followThreadsDomain.ID = existedFollowThreads.ID

	result, err := tu.followThreadsRepository.Delete(ctx, followThreadsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *followThreadsUsecase) GetStatus(ctx context.Context, threadID int, userID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if threadID <= 0 || userID <= 0 {
		return Domain{}, businesses.ErrIDResource
	}

	res, err := tu.followThreadsRepository.GetDuplicate(ctx, threadID, userID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}