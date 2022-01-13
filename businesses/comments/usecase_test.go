package comments_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/comments"
	_commentsMock "infion-BE/businesses/comments/mocks"
	"infion-BE/businesses/replies"
	_repliesMock "infion-BE/businesses/replies/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	commentsRepository _commentsMock.Repository
	commentsUsecase    comments.Usecase
	commentsDomain     comments.Domain
	repliesRepository	_repliesMock.Repository
	repliesDomain		replies.Domain
)

func TestMain(m *testing.M) {
	commentsUsecase = comments.NewCommentsUsecase(&commentsRepository, 2, &repliesRepository)
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
	commentsDomain = comments.Domain{
		ID:				1,
		ThreadID: 		1,
		UserID:			1,
		User:			"user1",
		UrlImg: 		"placeholder.jpg",
		Comment: 		"comment test.",
		Replies:		[]replies.Domain{repliesDomain},
		LikeCount:		1,
		ReplyCount:		1,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		commentsRepository.On("Store", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.Store(ctx, &commentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, commentsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		commentsRepository.On("Store", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.Store(ctx, &commentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.GetByID(ctx, commentsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, commentsDomain, result)
	})

	t.Run("GetByID | InValid commentsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := commentsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.GetByID(ctx, commentsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("GetByID | InValid Replies", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.GetByID(ctx, commentsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetComments(t *testing.T){
	t.Run("GetComments | Valid", func(t *testing.T) {
		commentsRepository.On("GetComments", mock.Anything).Return([]comments.Domain{commentsDomain}, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.GetComments(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []comments.Domain{commentsDomain}, result)
	})
	
	t.Run("GetComments | InValid", func(t *testing.T) {
		commentsRepository.On("GetComments", mock.Anything).Return([]comments.Domain{commentsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := commentsUsecase.GetComments(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})

	t.Run("GetComments | InValid Replies", func(t *testing.T) {
		commentsRepository.On("GetComments", mock.Anything).Return([]comments.Domain{commentsDomain}, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.GetComments(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetCommentsByThreadId(t *testing.T){
	t.Run("GetCommentsByThreadId | Valid", func(t *testing.T) {
		commentsRepository.On("GetCommentsByThreadID", mock.Anything, mock.AnythingOfType("int")).Return([]comments.Domain{commentsDomain}, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.GetCommentsByThreadID(ctx, commentsDomain.ThreadID)

		assert.Nil(t, err)
		assert.Equal(t, []comments.Domain{commentsDomain}, result)
	})

	t.Run("GetCommentsByThreadId | InValid", func(t *testing.T) {
		commentsRepository.On("GetCommentsByThreadID", mock.Anything, mock.AnythingOfType("int")).Return([]comments.Domain{commentsDomain}, businesses.ErrNotFound).Once()
		
		ctx := context.Background()
		_, err := commentsUsecase.GetCommentsByThreadID(ctx, commentsDomain.ThreadID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})

	t.Run("GetCommentsByThreadId | InValid Replies", func(t *testing.T) {
		commentsRepository.On("GetCommentsByThreadID", mock.Anything, mock.AnythingOfType("int")).Return([]comments.Domain{commentsDomain}, nil).Once()
		repliesRepository.On("GetRepliesByCommentID", mock.Anything, mock.AnythingOfType("int")).Return([]replies.Domain{repliesDomain}, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.GetCommentsByThreadID(ctx, commentsDomain.ThreadID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		commentsRepository.On("Update", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.Update(ctx, &commentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &commentsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.Update(ctx, &commentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		commentsRepository.On("Update", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.Update(ctx, &commentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		commentsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, nil).Once()

		ctx := context.Background()
		result, err := commentsUsecase.Delete(ctx, &commentsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &commentsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.Delete(ctx, &commentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		commentsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(commentsDomain, nil).Once()
		commentsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*comments.Domain")).Return(commentsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := commentsUsecase.Delete(ctx, &commentsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}