package threads_test

import (
	"context"
	"infion-BE/businesses"
	_commentsMock "infion-BE/businesses/comments/mocks"
	_followThreadsMock "infion-BE/businesses/followThreads/mocks"
	_likeThreadsMock "infion-BE/businesses/likeThreads/mocks"
	_reportsMock "infion-BE/businesses/reports/mocks"
	"infion-BE/businesses/threads"
	_threadsMock "infion-BE/businesses/threads/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	threadsRepository		_threadsMock.Repository
	threadsUsecase			threads.Usecase
	threadsDomain			threads.Domain
	likeThreadsRepository	_likeThreadsMock.Repository
	commentsRepository		_commentsMock.Repository
	followThreadsRepository	_followThreadsMock.Repository
	reportsRepository		_reportsMock.Repository
)

func TestMain(m *testing.M) {
	threadsUsecase = threads.NewThreadsUsecase(&threadsRepository, 2, &likeThreadsRepository, &commentsRepository, &followThreadsRepository, &reportsRepository)
	threadsDomain = threads.Domain{
		ID:				1,
		Title:			"test thread",
		Img:			"placeholder.jpg",
		Content:		"aaaaaaaaaaaaaa",
		Category:		"test",
		UserID:			1,
		User:			"user1",
		UrlImg: 		"placeholder.jpg",
		LikeCount:		1,
		CommentCount:	1,
		FollowerCount: 	1,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		threadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.Store(ctx, &threadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, threadsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		threadsRepository.On("Store", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.Store(ctx, &threadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, threadsDomain, result)
	})

	t.Run("GetByID | InValid threadsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrThreadsIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetByID(ctx, threadsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetThreads(t *testing.T){
	t.Run("GetThreads | Valid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetThreads(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []threads.Domain{threadsDomain}, result)
	})

	t.Run("GetThreads | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreads(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})

	
	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreads(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreads(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreads(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreads(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetThreadsByCategory(t *testing.T){
	t.Run("GetThreadsByCategory | Valid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.Nil(t, err)
		assert.Equal(t, []threads.Domain{threadsDomain}, result)
	})

	t.Run("GetThreadsByCategory | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
	
	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByCategory(ctx, "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetThreadsBySort(t *testing.T){
	t.Run("GetThreadsBySort | Valid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.Nil(t, err)
		assert.Equal(t, []threads.Domain{threadsDomain}, result)
	})

	t.Run("GetThreadsBySort | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySort", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySort(ctx, "sort")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetThreadsBySortCategory(t *testing.T){
	t.Run("GetThreadsBySortCategory | Valid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.Nil(t, err)
		assert.Equal(t, []threads.Domain{threadsDomain}, result)
	})

	t.Run("GetThreadsBySortCategory | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
	
	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsBySortCategory", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsBySortCategory(ctx, "sort", "category")

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetThreadsByUserID(t *testing.T){
	t.Run("GetThreadsByUserID | Valid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.Nil(t, err)
		assert.Equal(t, []threads.Domain{threadsDomain}, result)
	})

	t.Run("GetThreadsByUserID | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})

	
	t.Run("LikeCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		threadsRepository.On("GetThreadsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadsDomain}, nil).Once()
		likeThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		commentsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		followThreadsRepository.On("CountByThreadID", mock.Anything, mock.AnythingOfType("int")).Return(1, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.GetThreadsByUserID(ctx, threadsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.Update(ctx, &threadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &threadsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.Update(ctx, &threadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		threadsRepository.On("Update", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.Update(ctx, &threadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		threadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, nil).Once()

		ctx := context.Background()
		result, err := threadsUsecase.Delete(ctx, &threadsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &threadsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.Delete(ctx, &threadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(threadsDomain, nil).Once()
		threadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*threads.Domain")).Return(threadsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := threadsUsecase.Delete(ctx, &threadsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}