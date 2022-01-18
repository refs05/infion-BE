package controllers

import (
	"infion-BE/businesses/users"
	"infion-BE/controllers"
	"infion-BE/controllers/users/request"
	"infion-BE/controllers/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UseCase
}
func NewUserController(uc users.UseCase)*UserController{
	return &UserController{
		usecase: uc,
	}
}
func (controller *UserController) Login(c echo.Context)error{
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)

	}
	user,err := controller.usecase.Login(*userLogin.ToDomain(),ctx)
	return controllers.NewSuccessResponse(c,response.FromDomain(user))
}

func (controller *UserController) CreateNewUser(c echo.Context)error{


	ctx := c.Request().Context()
	var createUser request.CreateUser
	err := c.Bind(&createUser)

	// ctx := c.Request().Context()
	// var createUser request.CreateUser
	// err := c.Bind(&createUser)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)

	}
	user,err := controller.usecase.CreateNewUser(*createUser.ToDomain(),ctx)
	return controllers.NewSuccessResponse(c,response.FromDomain(user))
}

func (controller *UserController) FindById(c echo.Context)error{
	ctx := c.Request().Context()
	id,_:=strconv.Atoi(c.Param("id"))
	
	user,err := controller.usecase.FindById(id,ctx)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusBadGateway,err)
	}
	return controllers.NewSuccessResponse(c,response.FromDomain(user))


}

func (ctrl *UserController) GetLeaderboard(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := ctrl.usecase.GetLeaderboard(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	return controllers.NewSuccessResponse(c, response.NewResponseArray(users))
}