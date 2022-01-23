package users_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/users"
	_usersMock "infion-BE/businesses/users/mocks"
	"os"
	"testing"
	"time"

	_commentsMock "infion-BE/businesses/comments/mocks"
	_followThreadsMock "infion-BE/businesses/followThreads/mocks"
	_followUsersMock "infion-BE/businesses/followUsers/mocks"
	_threadsMock "infion-BE/businesses/threads/mocks"

	_middleware "infion-BE/app/middleware"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	usersRepository _usersMock.Repository
	usersUsecase    users.UseCase
	usersDomain     users.DomainUser
	threadsRepository		_threadsMock.Repository
	commentsRepository		_commentsMock.Repository
	followThreadsRepository	_followThreadsMock.Repository
	followUsersRepository	_followUsersMock.Repository
	jwt 					*_middleware.ConfigJWT
)

func TestMain(m *testing.M) {
	usersUsecase = users.NewUseCase(&usersRepository, 2, &commentsRepository, &threadsRepository, &followUsersRepository,  &followThreadsRepository, jwt)
	usersDomain = users.DomainUser{
		Id:				1,
		Username: "test user",
		Email 	: "test@email.com",
		Password : "password",
		RoleId	: 1,
		UrlImg	: "img.jpg",
		CommentCount: 0,
		ThreadCount: 0,
		LikeCount: 0,
		FollowerCount: 0,
		ThreadFollowerCount: 0,
		ThreadFollowingCount: 0,
		Rank          : 1,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
		Token:			"asadasdaa",
	}
	os.Exit(m.Run())
}

func TestLogin(t *testing.T){
	// t.Run("Login | Valid", func(t *testing.T) {
	// 	usersRepository.On("Login", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()
	// 	// jwt.On("")

	// 	testDomain := usersDomain
	// 	usersDomain.Password, _ = encrypt.Hash(usersDomain.Password)
	// 	fmt.Println(usersDomain.Password, testDomain.Password)
	// 	ctx := context.Background()
	// 	result, err := usersUsecase.Login(testDomain, ctx)

	// 	assert.Nil(t, err)
	// 	assert.NotEqual(t, usersDomain, result)
	// })

	t.Run("Login | InValid Password", func(t *testing.T) {
		usersRepository.On("Login", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()
		// jwt.On("")

		ctx := context.Background()
		_, err := usersUsecase.Login(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrWrongPass, err)
	})

	t.Run("Login | InValid Email 1", func(t *testing.T) {
		usersDomain.Email = ""
		ctx := context.Background()
		_, err := usersUsecase.Login(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrUsernamePasswordNotFound, err)
	})

	t.Run("Login | InValid Email 2", func(t *testing.T) {
		usersDomain.Email = "testemail.com"
		ctx := context.Background()
		_, err := usersUsecase.Login(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrWrongFormat, err)
	})

	t.Run("Login | InValid Password", func(t *testing.T) {
		usersDomain.Email = "test@email.com"
		usersDomain.Password = ""
		ctx := context.Background()
		_, err := usersUsecase.Login(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrUsernamePasswordNotFound, err)
	})

	t.Run("Login | InValid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, businesses.ErrInternalServer).Once()
		usersDomain.Email = "test@email.com"
		usersDomain.Password = "password"
		ctx := context.Background()
		_, err := usersUsecase.Login(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestCreateNewUser(t *testing.T){
	t.Run("CreateNewUser | InValid Duplicate", func(t *testing.T) {
		usersRepository.On("GetUsername", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()
		usersRepository.On("CreateNewUser", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()

		ctx := context.Background()
		_, err := usersUsecase.CreateNewUser(usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrDuplicateUsername, err)
	})

	t.Run("CreateNewUser | InValid empty email", func(t *testing.T) {
		usersDomain2 := usersDomain
		usersDomain2.Email = ""
		ctx := context.Background()
		_, err := usersUsecase.CreateNewUser(usersDomain2, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrUsernamePasswordNotFound, err)
	})

	t.Run("CreateNewUser | InValid empty password", func(t *testing.T) {
		usersDomain2 := usersDomain
		usersDomain2.Password = ""
		ctx := context.Background()
		_, err := usersUsecase.CreateNewUser(usersDomain2, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrUsernamePasswordNotFound, err)
	})

	t.Run("CreateNewUser | Valid", func(t *testing.T) {
		usersRepository.On("GetUsername", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, gorm.ErrRecordNotFound).Once()
		usersRepository.On("CreateNewUser", mock.AnythingOfType("users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()

		ctx := context.Background()
		result, err := usersUsecase.CreateNewUser(usersDomain, ctx)

		assert.Nil(t, err)
		assert.Equal(t, usersDomain, result)
	})
}

func TestFindById(t *testing.T){
	t.Run("FindById | Valid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()
		usersRepository.On("FindById", mock.AnythingOfType("int"), mock.Anything).Return(usersDomain, nil).Once()

		ctx := context.Background()
		result, err := usersUsecase.FindById(int(usersDomain.Id), ctx)

		assert.Nil(t, err)
		assert.Equal(t, usersDomain, result)
	})

	t.Run("FindById Leaderboard | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.FindById(int(usersDomain.Id), ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FindById | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()
		usersRepository.On("FindById", mock.AnythingOfType("int"), mock.Anything).Return(usersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.FindById(int(usersDomain.Id), ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		usersRepository.On("FindById", mock.AnythingOfType("int"), mock.Anything).Return(usersDomain, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()

		ctx := context.Background()
		result, err := usersUsecase.Update(&usersDomain, ctx)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Find By Id | InValid", func(t *testing.T) {
		usersRepository.On("FindById", mock.AnythingOfType("int"), mock.Anything).Return(usersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.Update(&usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("FindById", mock.AnythingOfType("int"), mock.Anything).Return(usersDomain, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain,  businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.Update(&usersDomain, ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetLeaderboards(t *testing.T){
	t.Run("GetLeaderboard | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("ThreadLikeCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("CommentCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("ThreadCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("FollowerCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("ThreadFollowerCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("ThreadFollowingCount | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := usersUsecase.GetLeaderboard(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Get Leaderboards | Valid", func(t *testing.T) {
		usersRepository.On("GetLeaderboard", mock.Anything).Return([]users.DomainUser{usersDomain}, nil).Once()
		threadsRepository.On("GetThreadLikeCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		commentsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followUsersRepository.On("CountByFollowedID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		threadsRepository.On("GetThreadFollowerCountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		followThreadsRepository.On("CountByUserID", mock.Anything, mock.AnythingOfType("uint")).Return(0, nil).Once()
		usersRepository.On("Update", mock.AnythingOfType("*users.DomainUser"), mock.Anything).Return(usersDomain, nil).Once()

		ctx := context.Background()
		result, err := usersUsecase.GetLeaderboard(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []users.DomainUser{usersDomain}, result)
	})
}