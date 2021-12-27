package followThreads_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/followThreads"
	_followThreadsMock "infion-BE/businesses/followThreads/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	followThreadsRepository _followThreadsMock.Repository
	followThreadsUsecase followThreads.Usecase
	followThreadsDomain followThreads.Domain
)

func TestMain(m *testing.M) {
	followThreadsUsecase = followThreads.NewFollowThreadsUsecase(&followThreadsRepository, 2)
	followThreadsDomain = followThreads.Domain{
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
		followThreadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := followThreadsUsecase.Store(ctx, &followThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, followThreadsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		followThreadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.Store(ctx, &followThreadsDomain)

		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := followThreadsUsecase.GetByID(ctx, followThreadsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, followThreadsDomain, result)
	})

	t.Run("GetByID | InValid followThreadsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := followThreadsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.GetByID(ctx, followThreadsDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, nil).Once()
		followThreadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := followThreadsUsecase.Update(ctx, &followThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &followThreadsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.Update(ctx, &followThreadsDomain)

		assert.NotNil(t, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, nil).Once()
		followThreadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.Update(ctx, &followThreadsDomain)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, nil).Once()
		followThreadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, nil).Once()

		ctx := context.Background()
		result, err := followThreadsUsecase.Delete(ctx, &followThreadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &followThreadsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.Delete(ctx, &followThreadsDomain)

		assert.NotNil(t, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		followThreadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followThreadsDomain, nil).Once()
		followThreadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*followThreads.Domain")).Return(followThreadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followThreadsUsecase.Delete(ctx, &followThreadsDomain)

		assert.NotNil(t, err)
	})
}