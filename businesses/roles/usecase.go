package roles

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type rolesUsecase struct {
	rolesRepository  Repository
	contextTimeout  time.Duration
}

func NewRolesUsecase(tr Repository, timeout time.Duration) Usecase {
	return &rolesUsecase{
		rolesRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *rolesUsecase) Store(ctx context.Context, rolesDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.rolesRepository.Store(ctx, rolesDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *rolesUsecase) GetByID(ctx context.Context, rolesId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if rolesId <= 0 {
		return Domain{}, businesses.ErrRolesIDResource
	}
	res, err := tu.rolesRepository.GetByID(ctx, rolesId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *rolesUsecase) Update(ctx context.Context, rolesDomain *Domain) (*Domain, error) {
	existedRoles, err := tu.rolesRepository.GetByID(ctx, rolesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	rolesDomain.ID = existedRoles.ID

	result, err := tu.rolesRepository.Update(ctx, rolesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *rolesUsecase) Delete(ctx context.Context, rolesDomain *Domain) (*Domain, error) {
	existedRoles, err := tu.rolesRepository.GetByID(ctx, rolesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	rolesDomain.ID = existedRoles.ID

	result, err := tu.rolesRepository.Delete(ctx, rolesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}