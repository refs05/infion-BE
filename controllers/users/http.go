package controllers

import (
	"infion-BE/bussiness/users"
	"infion-BE/controllers"
	"infion-BE/controllers/users/request"
	"infion-BE/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}
func NewUserController(uc users.UserUseCaseInterface)*UserController{
	return &UserController{
		usecase: uc,
	}
}
func (controller *UserController) Login(c echo.Context)error{
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)

	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"error binding",err)

	}
	user,err := controller.usecase.Login(*userLogin.ToDomain(),ctx)
	return controllers.SuccessResponse(c,response.FromDomain(user))
}

func (controller *UserController) CreateNewUser(c echo.Context)error{


	ctx := c.Request().Context()
	var createUser request.CreateUser
	err := c.Bind(&createUser)

	// ctx := c.Request().Context()
	// var createUser request.CreateUser
	// err := c.Bind(&createUser)

	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"error binding",err)

	}
	user,err := controller.usecase.CreateNewUser(*createUser.ToDomain(),ctx)
	return controllers.SuccessResponse(c,response.FromDomain(user))
}