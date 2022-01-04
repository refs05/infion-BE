package replies_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/replies"
	_repliesMock "infion-BE/businesses/replies/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repliesRepository _repliesMock.Repository
	repliesUsecase    replies.Usecase
	repliesDomain     replies.Domain
)

func TestMain(m *testing.M) {
	repliesUsecase = replies.NewRepliesUsecase(&repliesRepository, 2)
	repliesDomain = replies.Domain{
		ID:        1,
		CommentID: 1,
		UserID:    1,
		User:      "user1",
		UrlImg:    "placeholder.jpg",
		Reply:     "comment test.",
		LikeCount: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T) {
	t.Run("Store | Valid", func(t *testing.T) {
		repliesRepository.On("Store", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.Store(ctx, &repliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, repliesDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		repliesRepository.On("Store", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Store(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.GetByID(ctx, repliesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, repliesDomain, result)
	})

	t.Run("GetByID | InValid repliesId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := repliesUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetByID(ctx, repliesDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetReplies(t *testing.T) {
	t.Run("GetReplies | Valid", func(t *testing.T) {
		repliesRepository.On("GetReplies", mock.Anything).Return([]replies.Domain{repliesDomain}, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.GetReplies(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []replies.Domain{repliesDomain}, result)
	})

	t.Run("GetReplies | InValid", func(t *testing.T) {
		repliesRepository.On("GetReplies", mock.Anything).Return([]replies.Domain{repliesDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetReplies(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.Update(ctx, &repliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &repliesDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Update(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Update(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		repliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.Delete(ctx, &repliesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &repliesDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Delete(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		repliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Delete(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}
