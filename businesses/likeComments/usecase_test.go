package likeComments_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeComments"
	_likeCommentsMock "infion-BE/businesses/likeComments/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	likeCommentsRepository _likeCommentsMock.Repository
	likeCommentsUsecase likeComments.Usecase
	likeCommentsDomain likeComments.Domain
)

func TestMain(m *testing.M) {
	likeCommentsUsecase = likeComments.NewLikeCommentsUsecase(&likeCommentsRepository, 2)
	likeCommentsDomain = likeComments.Domain{
		ID:         1,
		CommentID: 	1,
		UserID: 	1,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		likeCommentsRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeCommentsDomain, businesses.ErrDuplicateData).Once()
		likeCommentsRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeCommentsUsecase.Store(ctx, &likeCommentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, likeCommentsDomain, result)
	})

	t.Run("Duplicate | Status Update", func(t *testing.T) {
		likeCommentsRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()
		likeCommentsRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeCommentsUsecase.Store(ctx, &likeCommentsDomain)

		assert.Nil(t, err)
		assert.NotEqual(t, likeCommentsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		likeCommentsRepository.On("GetDuplicate", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(likeCommentsDomain, businesses.ErrDuplicateData).Once()
		likeCommentsRepository.On("Store", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.Store(ctx, &likeCommentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeCommentsUsecase.GetByID(ctx, likeCommentsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, likeCommentsDomain, result)
	})

	t.Run("GetByID | InValid likeCommentsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := likeCommentsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.GetByID(ctx, likeCommentsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()
		likeCommentsRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeCommentsUsecase.Update(ctx, &likeCommentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeCommentsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.Update(ctx, &likeCommentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()
		likeCommentsRepository.On("Update", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.Update(ctx, &likeCommentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()
		likeCommentsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, nil).Once()

		ctx := context.Background()
		result, err := likeCommentsUsecase.Delete(ctx, &likeCommentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &likeCommentsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.Delete(ctx, &likeCommentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		likeCommentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(likeCommentsDomain, nil).Once()
		likeCommentsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeComments.Domain")).Return(likeCommentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := likeCommentsUsecase.Delete(ctx, &likeCommentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}