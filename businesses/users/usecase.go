package users

import (
	"context"

	_middleware "infion-BE/app/middleware"
	"infion-BE/businesses"
	"infion-BE/businesses/comments"
	"infion-BE/businesses/followThreads"
	"infion-BE/businesses/followUsers"
	"infion-BE/businesses/threads"
	"infion-BE/drivers/helpers/encrypt"

	"regexp"
	"time"

	"gorm.io/gorm"
)

type UserUseCase struct {
	repo Repository
	ctx  time.Duration
	commentsRepository		comments.Repository
	threadsRepository		threads.Repository
	followUsersRepository	followUsers.Repository
	followThreadsRepository	followThreads.Repository
	jwt 					*_middleware.ConfigJWT
}

func NewUseCase(UserRepo Repository,contextTimeout time.Duration, cr comments.Repository, tr threads.Repository, fur followUsers.Repository, ftr followThreads.Repository, configJWT *_middleware.ConfigJWT) UseCase{
	return &UserUseCase{
		repo: UserRepo,
		ctx: contextTimeout,
		commentsRepository: cr,
		threadsRepository: tr,
		followUsersRepository: fur,
		followThreadsRepository: ftr,
    	jwt: configJWT,
	}
}


func (usecase *UserUseCase) Login(domain DomainUser, ctx context.Context )(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, businesses.ErrUsernamePasswordNotFound
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
   	if !emailRegex.MatchString(domain.Email){
		return DomainUser{}, businesses.ErrWrongFormat
	}

	if domain.Password == ""{
		return DomainUser{},businesses.ErrUsernamePasswordNotFound
	}

	// password := domain.Password
	// hash,_:= encrypt.Hash(domain.Password)

	// if !encrypt.ValidateHash(password,hash) {
	// 	return DomainUser{}, errors.New("Wrong Password")
	// }

	user,err := usecase.repo.Login(domain,ctx)
	if err !=nil{
		return DomainUser{},err
	}

	if !encrypt.ValidateHash(domain.Password, user.Password) {
		return DomainUser{},businesses.ErrWrongPass
	}

	user.Token = usecase.jwt.GenererateToken(user.Id)
	return user,nil
}

func (usecase *UserUseCase) CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, businesses.ErrUsernamePasswordNotFound
	}
	if domain.Password == ""{
		return DomainUser{},businesses.ErrUsernamePasswordNotFound
	}

	existedUser,err := usecase.repo.GetUsername(domain,ctx)
	if err == gorm.ErrRecordNotFound {
		domain.Password,err = encrypt.Hash(domain.Password)
		if err != nil{
			return DomainUser{},businesses.ErrInternalServer
		}
		user,err := usecase.repo.CreateNewUser(domain,ctx)
		if err != nil{
			return DomainUser{},err
		}
		return user,nil
	} else {
		return existedUser,businesses.ErrDuplicateUsername
	}
}

func (usecase *UserUseCase)FindById(userId int,ctx context.Context)(DomainUser,error){
	_, err := usecase.GetLeaderboard(ctx)
	if err != nil{
		return DomainUser{},err
	}
	rec,err := usecase.repo.FindById(userId,ctx)
	if err != nil{
		return DomainUser{},err
	}
	return rec, nil
}

func (usecase *UserUseCase) Update(userDomain *DomainUser, ctx context.Context) (*DomainUser, error) {
	existedComments, err := usecase.repo.FindById(int(userDomain.Id),ctx)
	if err != nil {
		return &DomainUser{}, err
	}
	userDomain.Id = existedComments.Id

	result, err := usecase.repo.Update(userDomain, ctx)
	if err != nil {
		return &DomainUser{}, err
	}

	return &result, nil
}

func (usecase *UserUseCase)GetLeaderboard(ctx context.Context)([]DomainUser,error){
	result, err := usecase.repo.GetLeaderboard(ctx)
	if err != nil {
		return []DomainUser{}, err
	}

	for i := range result {
		result[i].LikeCount, err = usecase.threadsRepository.GetThreadLikeCountByUserID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].ThreadCount, err = usecase.commentsRepository.CountByUserID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].ThreadCount, err = usecase.threadsRepository.CountByUserID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].FollowerCount, err = usecase.followUsersRepository.CountByFollowedID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].ThreadFollowerCount, err = usecase.threadsRepository.GetThreadFollowerCountByUserID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].ThreadFollowingCount, err = usecase.followThreadsRepository.CountByUserID(ctx, result[i].Id)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	for i := range result {
		result[i].Rank = i + 1
	}

	for i := range result {
		_, err = usecase.repo.Update(&result[i], ctx)
		if err != nil {
			return []DomainUser{}, err
		}
	}

	return result, nil
}
