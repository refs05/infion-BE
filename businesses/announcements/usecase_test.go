package announcements_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/announcements"
	_announcementsMock "infion-BE/businesses/announcements/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	announcementsRepository _announcementsMock.Repository
	announcementsUsecase    announcements.Usecase
	announcementsDomain     announcements.Domain
)

func TestMain(m *testing.M) {
	announcementsUsecase = announcements.NewAnnouncementsUsecase(&announcementsRepository, 2)
	announcementsDomain = announcements.Domain{
		ID:				1,
		UserID:			1,
		Announcer:		"mod1",
		Message:		"test message",
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		announcementsRepository.On("Store", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.Store(ctx, &announcementsDomain)

		assert.Nil(t, err)
		assert.Equal(t, announcementsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		announcementsRepository.On("Store", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.Store(ctx, &announcementsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.GetByID(ctx, announcementsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, announcementsDomain, result)
	})

	t.Run("GetByID | InValid announcementsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := announcementsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.GetByID(ctx, announcementsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetAnnouncements(t *testing.T){
	t.Run("GetAnnouncements | Valid", func(t *testing.T) {
		announcementsRepository.On("GetAnnouncements", mock.Anything).Return([]announcements.Domain{announcementsDomain}, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.GetAnnouncements(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []announcements.Domain{announcementsDomain}, result)
	})

	t.Run("GetAnnouncements | InValid", func(t *testing.T) {
		announcementsRepository.On("GetAnnouncements", mock.Anything).Return([]announcements.Domain{announcementsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.GetAnnouncements(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
}

func TestGetAnnouncementsByUserID(t *testing.T){
	t.Run("GetAnnouncements | Valid", func(t *testing.T) {
		announcementsRepository.On("GetAnnouncementsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]announcements.Domain{announcementsDomain}, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.GetAnnouncementsByUserID(ctx, announcementsDomain.UserID)

		assert.Nil(t, err)
		assert.Equal(t, []announcements.Domain{announcementsDomain}, result)
	})

	t.Run("GetAnnouncements | InValid", func(t *testing.T) {
		announcementsRepository.On("GetAnnouncementsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]announcements.Domain{announcementsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.GetAnnouncementsByUserID(ctx, announcementsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, nil).Once()
		announcementsRepository.On("Update", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.Update(ctx, &announcementsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &announcementsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.Update(ctx, &announcementsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, nil).Once()
		announcementsRepository.On("Update", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.Update(ctx, &announcementsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, nil).Once()
		announcementsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, nil).Once()

		ctx := context.Background()
		result, err := announcementsUsecase.Delete(ctx, &announcementsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &announcementsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.Delete(ctx, &announcementsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		announcementsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(announcementsDomain, nil).Once()
		announcementsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*announcements.Domain")).Return(announcementsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := announcementsUsecase.Delete(ctx, &announcementsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}