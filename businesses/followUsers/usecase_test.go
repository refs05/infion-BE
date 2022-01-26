package followUsers_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/followUsers"
	_followUsersMock "infion-BE/businesses/followUsers/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	followUsersRepository _followUsersMock.Repository
	followUsersUsecase followUsers.Usecase
	followUsersDomain followUsers.Domain
)

func TestMain(m *testing.M) {
	followUsersUsecase = followUsers.NewFollowUsersUsecase(&followUsersRepository, 2)
	followUsersDomain = followUsers.Domain{
		ID:         1,
		FollowedID: 1,
		FollowerID:	1,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		followUsersRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrDuplicateData).Once()
		followUsersRepository.On("Store", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.Store(ctx, &followUsersDomain)

		assert.Nil(t, err)
		assert.Equal(t, followUsersDomain, result)
	})

	t.Run("Duplicate | Status Update", func(t *testing.T) {
		followUsersRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()
		followUsersRepository.On("Update", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.Store(ctx, &followUsersDomain)

		assert.Nil(t, err)
		assert.NotEqual(t, followUsersDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		followUsersRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrDuplicateData).Once()
		followUsersRepository.On("Store", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.Store(ctx, &followUsersDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.GetByID(ctx, followUsersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, followUsersDomain, result)
	})

	t.Run("GetByID | InValid followUsersId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := followUsersUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.GetByID(ctx, followUsersDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()
		followUsersRepository.On("Update", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.Update(ctx, &followUsersDomain)

		assert.Nil(t, err)
		assert.Equal(t, &followUsersDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.Update(ctx, &followUsersDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()
		followUsersRepository.On("Update", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.Update(ctx, &followUsersDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()
		followUsersRepository.On("Delete", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.Delete(ctx, &followUsersDomain)

		assert.Nil(t, err)
		assert.Equal(t, &followUsersDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.Delete(ctx, &followUsersDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		followUsersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()
		followUsersRepository.On("Delete", mock.Anything, mock.AnythingOfType("*followUsers.Domain")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.Delete(ctx, &followUsersDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetStatus(t *testing.T){
	t.Run("GetStatus | Valid", func(t *testing.T) {
		followUsersRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(followUsersDomain, nil).Once()

		ctx := context.Background()
		result, err := followUsersUsecase.GetStatus(ctx, followUsersDomain.FollowedID, followUsersDomain.FollowerID)

		assert.Nil(t, err)
		assert.Equal(t, followUsersDomain, result)
	})

	t.Run("GetStatus | InValid ID <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := followUsersUsecase.GetStatus(ctx, 0, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetStatus | InValid", func(t *testing.T) {
		followUsersRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(followUsersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := followUsersUsecase.GetStatus(ctx, followUsersDomain.FollowedID, followUsersDomain.FollowerID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}