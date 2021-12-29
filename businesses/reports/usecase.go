package reports

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type reportsUsecase struct {
	reportsRepository  Repository
	contextTimeout  time.Duration
}

func NewReportsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &reportsUsecase{
		reportsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *reportsUsecase) Store(ctx context.Context, reportsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.reportsRepository.Store(ctx, reportsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *reportsUsecase) GetByID(ctx context.Context, reportsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if reportsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.reportsRepository.GetByID(ctx, reportsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *reportsUsecase) GetReports(ctx context.Context) ([]Domain, error) {
	result, err := tu.reportsRepository.GetReports(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (tu *reportsUsecase) Update(ctx context.Context, reportsDomain *Domain) (*Domain, error) {
	existedReports, err := tu.reportsRepository.GetByID(ctx, reportsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	reportsDomain.ID = existedReports.ID

	result, err := tu.reportsRepository.Update(ctx, reportsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *reportsUsecase) Delete(ctx context.Context, reportsDomain *Domain) (*Domain, error) {
	existedReports, err := tu.reportsRepository.GetByID(ctx, reportsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	reportsDomain.ID = existedReports.ID

	result, err := tu.reportsRepository.Delete(ctx, reportsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}