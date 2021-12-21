package users

import (
	"context"
	"errors"
	
	// "infion-BE/drivers/helpers/encrypt"
	"regexp"
	// "strings"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration

}
func NewUseCase(UserRepo UserRepoInterface,contextTimeout time.Duration)UserUseCaseInterface{
	return &UserUseCase{
		repo: UserRepo,
		ctx: contextTimeout,
		
	}
}


func (usecase *UserUseCase) Login(domain DomainUser, ctx context.Context )(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, errors.New("Email Empty")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

   	if emailRegex.MatchString(domain.Email) != true{
		return DomainUser{}, errors.New("Wrong Email Format")
	   } 
	if domain.Password == ""{
		return DomainUser{},errors.New("Password Empty")
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
	return user,nil
}

func (usecase *UserUseCase) CreateNewUser(domain DomainUser, ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, errors.New("Email Empty")
	}
	if domain.Password == ""{
		return DomainUser{},errors.New("Password Empty")
	}
// existedUser,err := usecase.repo.GetUsername(domain,ctx)

// if err != nil {

// 	if !strings.Contains(err.Error(), "not found") {
// 		return DomainUser{}, err
// 	}
// }
// if existedUser !=(DomainUser{}){
// 	return  DomainUser{},errors.New("Duplicate data")
// }

// domain.Password,err = encrypt.Hash(domain.Password)
// if err != nil{
// 	return DomainUser{},errors.New("something gone wrong, contact administrator")
// }
	user,err := usecase.repo.CreateNewUser(domain,ctx)
	if err != nil{
		return DomainUser{},err
	}
return user,nil
}

