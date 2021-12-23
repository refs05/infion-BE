package users

import (
	"context"

	"infion-BE/businesses"

	"infion-BE/drivers/helpers/encrypt"

	"regexp"
	// "strings"
	"time"

	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type UserUseCase struct {
	repo Repository
	ctx  time.Duration

}
func NewUseCase(UserRepo Repository,contextTimeout time.Duration) UseCase{
	return &UserUseCase{
		repo: UserRepo,
		ctx: contextTimeout,
		
	}
}


func (usecase *UserUseCase) Login(domain DomainUser, ctx context.Context )(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, businesses.ErrUsernamePasswordNotFound
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

   	if emailRegex.MatchString(domain.Email) != true{
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
	
	
}else{
return existedUser,businesses.ErrDuplicateData
}
}

func (usecase *UserUseCase)FindById(userId int,ctx context.Context)(DomainUser,error){
	rec,err := usecase.repo.FindById(userId,ctx)
 if err != nil{
	 return DomainUser{},err
 }
 return rec, nil
}


