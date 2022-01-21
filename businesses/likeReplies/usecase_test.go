package likeReplies_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeReplies"
	_likeRepliesMock "infion-BE/businesses/likeReplies/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	likeRepliesRepository _likeRepliesMock.Repository
	likeRepliesUsecase likeReplies.Usecase
	likeRepliesDomain likeReplies.Domain
)

func TestMain(m *testing.M) {
	likeRepliesUsecase = likeReplies.NewLikeRepliesUsecase(&likeRepliesRepository, 2)
	likeRepliesDomain = likeReplies.Domain{
		ID:         1,
		ReplyID: 	1,
		UserID: 	1,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		likeRepliesRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeRepliesDomain, businesses.ErrDuplicateData).Once()
		likeRepliesRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, nil).Once()

		ctx := context.Background()
		result, err := likeRepliesUsecase.Store(ctx, &likeRepliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, likeRepliesDomain, result)
	})

	t.Run("Duplicate | InValid", func(t *testing.T) {
		likeRepliesRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Store(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		likeRepliesRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeRepliesDomain, businesses.ErrDuplicateData).Once()
		likeRepliesRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Store(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()

		ctx := context.Background()
		result, err := likeRepliesUsecase.GetByID(ctx, likeRepliesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, likeRepliesDomain, result)
	})

	t.Run("GetByID | InValid likeRepliesId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := likeRepliesUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.GetByID(ctx, likeRepliesDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()
		likeRepliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, nil).Once()

		ctx := context.Background()
		result, err := likeRepliesUsecase.Update(ctx, &likeRepliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeRepliesDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Update(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()
		likeRepliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Update(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()
		likeRepliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, nil).Once()

		ctx := context.Background()
		result, err := likeRepliesUsecase.Delete(ctx, &likeRepliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeRepliesDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Delete(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		likeRepliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeRepliesDomain, nil).Once()
		likeRepliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeRepliesUsecase.Delete(ctx, &likeRepliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}