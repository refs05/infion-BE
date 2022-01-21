package replies_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/likeReplies"
	_likeRepliesMock "infion-BE/businesses/likeReplies/mocks"
	"infion-BE/businesses/replies"
	_repliesMock "infion-BE/businesses/replies/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repliesRepository		_repliesMock.Repository
	repliesUsecase			replies.Usecase
	repliesDomain			replies.Domain
	likeRepliesRepository	_likeRepliesMock.Repository
	likeRepliesDomain		likeReplies.Domain
)

func TestMain(m *testing.M) {
	repliesUsecase = replies.NewRepliesUsecase(&repliesRepository, 2, &likeRepliesRepository)
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
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

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

	t.Run("LikeCount | InValid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetByID(ctx, repliesDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		repliesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(repliesDomain, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetByID(ctx, repliesDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetReplies(t *testing.T) {
	t.Run("GetReplies | Valid", func(t *testing.T) {
		repliesRepository.On("GetReplies", mock.Anything).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

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

	t.Run("LikeCount | InValid", func(t *testing.T) {
		repliesRepository.On("GetReplies", mock.Anything).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetReplies(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		repliesRepository.On("GetReplies", mock.Anything).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetReplies(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetRepliesByCommentId(t *testing.T){
	t.Run("GetRepliesByCommentId | Valid", func(t *testing.T) {
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, nil).Once()

		ctx := context.Background()
		result, err := repliesUsecase.GetRepliesByCommentID(ctx, repliesDomain.CommentID)

		assert.Nil(t, err)
		assert.Equal(t, []replies.Domain{repliesDomain}, result)
	})

	t.Run("GetRepliesByCommentId | InValid", func(t *testing.T) {
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetRepliesByCommentID(ctx, repliesDomain.CommentID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})

	t.Run("LikeCount | InValid", func(t *testing.T) {
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetRepliesByCommentID(ctx, repliesDomain.CommentID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()
		likeRepliesRepository.On("CountByReplyID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		repliesRepository.On("Update", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.GetRepliesByCommentID(ctx, repliesDomain.CommentID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
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
		likeRepliesRepository.On("GetLikeRepliesByReplyID", mock.Anything, mock.AnythingOfType("int")).Return([]likeReplies.Domain{likeRepliesDomain}, nil).Once()
		likeRepliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, nil).Once()
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
		likeRepliesRepository.On("GetLikeRepliesByReplyID", mock.Anything, mock.AnythingOfType("int")).Return([]likeReplies.Domain{likeRepliesDomain}, nil).Once()
		likeRepliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*likeReplies.Domain")).Return(likeRepliesDomain, nil).Once()
		repliesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*replies.Domain")).Return(repliesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := repliesUsecase.Delete(ctx, &repliesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}
