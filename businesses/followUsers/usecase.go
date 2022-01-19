package followUsers

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type followUsersUsecase struct {
	followUsersRepository  Repository
	contextTimeout  time.Duration
}

func NewFollowUsersUsecase(tr Repository, timeout time.Duration) Usecase {
	return &followUsersUsecase{
		followUsersRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *followUsersUsecase) Store(ctx context.Context, followUsersDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.followUsersRepository.Store(ctx, followUsersDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *followUsersUsecase) GetByID(ctx context.Context, followUsersId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if followUsersId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.followUsersRepository.GetByID(ctx, followUsersId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *followUsersUsecase) Update(ctx context.Context, followUsersDomain *Domain) (*Domain, error) {
	existedFollowUsers, err := tu.followUsersRepository.GetByID(ctx, followUsersDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	followUsersDomain.ID = existedFollowUsers.ID

	result, err := tu.followUsersRepository.Update(ctx, followUsersDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *followUsersUsecase) Delete(ctx context.Context, followUsersDomain *Domain) (*Domain, error) {
	existedFollowUsers, err := tu.followUsersRepository.GetByID(ctx, followUsersDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	followUsersDomain.ID = existedFollowUsers.ID

	result, err := tu.followUsersRepository.Delete(ctx, followUsersDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}