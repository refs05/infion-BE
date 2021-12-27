package likeThreads_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeThreads"
	_likeThreadsMock "infion-BE/businesses/likeThreads/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	likeThreadsRepository _likeThreadsMock.Repository
	likeThreadsUsecase likeThreads.Usecase
	likeThreadsDomain likeThreads.Domain
)

func TestMain(m *testing.M) {
	likeThreadsUsecase = likeThreads.NewLikeThreadsUsecase(&likeThreadsRepository, 2)
	likeThreadsDomain = likeThreads.Domain{
		ID:         1,
		ThreadID: 	1,
		UserID: 	1,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		likeThreadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeThreadsUsecase.Store(ctx, &likeThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, likeThreadsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		likeThreadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.Store(ctx, &likeThreadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeThreadsUsecase.GetByID(ctx, likeThreadsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, likeThreadsDomain, result)
	})

	t.Run("GetByID | InValid likeThreadsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := likeThreadsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.GetByID(ctx, likeThreadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, nil).Once()
		likeThreadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeThreadsUsecase.Update(ctx, &likeThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeThreadsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.Update(ctx, &likeThreadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, nil).Once()
		likeThreadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.Update(ctx, &likeThreadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, nil).Once()
		likeThreadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeThreadsUsecase.Delete(ctx, &likeThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeThreadsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.Delete(ctx, &likeThreadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		likeThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeThreadsDomain, nil).Once()
		likeThreadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeThreads.Domain")).Return(likeThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeThreadsUsecase.Delete(ctx, &likeThreadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}