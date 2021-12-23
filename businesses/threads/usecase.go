package threads

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type threadsUsecase struct {
	threadsRepository  Repository
	contextTimeout  time.Duration
}

func NewThreadsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &threadsUsecase{
		threadsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *threadsUsecase) GetThreads(ctx context.Context) ([]Domain, error) {
	result, err := tu.threadsRepository.GetThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (tu *threadsUsecase) Store(ctx context.Context, threadsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	// existedThreads, err := tu.threadsRepository.GetByTitle(ctx, threadsDomain.Title)
	// if err != nil {
	// 	if !strings.Contains(err.Error(), "not found") {
	// 		return Domain{}, err
	// 	}
	// }
	// if existedThreads != (Domain{}) {
	// 	return Domain{}, businesses.ErrDuplicateData
	// }

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

	return res, nil
}

// func (tu *threadsUsecase) GetByTitle(ctx context.Context, threadsTitle string) (Domain, error) {
// 	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
// 	defer cancel()

// 	if strings.TrimSpace(threadsTitle) == "" {
// 		return Domain{}, businesses.ErrThreadsTitleResource
// 	}
// 	res, err := tu.threadsRepository.GetByTitle(ctx, threadsTitle)
// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return res, nil
// }
// func (tu *threadsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
// 	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
// 	defer cancel()

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if perpage <= 0 {
// 		perpage = 25
// 	}

// 	res, total, err := tu.threadsRepository.Fetch(ctx, page, perpage)
// 	if err != nil {
// 		return []Domain{}, 0, err
// 	}

// 	return res, total, nil
// }

func (tu *threadsUsecase) Update(ctx context.Context, threadsDomain *Domain) (*Domain, error) {
	existedThreads, err := tu.threadsRepository.GetByID(ctx, threadsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	threadsDomain.ID = existedThreads.ID

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