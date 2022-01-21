package announcements

import (
	"context"
	"infion-BE/businesses"
	"time"
)

type announcementsUsecase struct {
	announcementsRepository  Repository
	contextTimeout  time.Duration
}

func NewAnnouncementsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &announcementsUsecase{
		announcementsRepository:  tr,
		contextTimeout:  timeout,
	}
}

func (tu *announcementsUsecase) Store(ctx context.Context, announcementsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	result, err := tu.announcementsRepository.Store(ctx, announcementsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *announcementsUsecase) GetByID(ctx context.Context, announcementsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if announcementsId <= 0 {
		return Domain{}, businesses.ErrIDResource
	}
	res, err := tu.announcementsRepository.GetByID(ctx, announcementsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *announcementsUsecase) GetAnnouncements(ctx context.Context) ([]Domain, error) {
	result, err := tu.announcementsRepository.GetAnnouncements(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (tu *announcementsUsecase) GetAnnouncementsByUserID(ctx context.Context, userID int) ([]Domain, error) {
	result, err := tu.announcementsRepository.GetAnnouncementsByUserID(ctx, userID)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (tu *announcementsUsecase) Update(ctx context.Context, announcementsDomain *Domain) (*Domain, error) {
	existedAnnouncements, err := tu.announcementsRepository.GetByID(ctx, announcementsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	announcementsDomain.ID = existedAnnouncements.ID

	result, err := tu.announcementsRepository.Update(ctx, announcementsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (tu *announcementsUsecase) Delete(ctx context.Context, announcementsDomain *Domain) (*Domain, error) {
	existedAnnouncements, err := tu.announcementsRepository.GetByID(ctx, announcementsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	announcementsDomain.ID = existedAnnouncements.ID

	result, err := tu.announcementsRepository.Delete(ctx, announcementsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}