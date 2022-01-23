package controllers

import (
	"infion-BE/businesses/followUser"
	"infion-BE/controllers"
	"infion-BE/controllers/followUser/response"
	"infion-BE/controllers/followUser/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FollowUserController struct {
	usecase followUser.UseCase
}

func NewFollowUserController(uc followUser.UseCase)*FollowUserController{
	return &FollowUserController{
		usecase: uc,
	}
}
func (controller *FollowUserController) Follow(c echo.Context) error{
	ctx := c.Request().Context()

	var follow request.FollowUser

	err := c.Bind(&follow)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}
	followUserr,err := controller.usecase.Follow(*follow.ToDomain(),ctx)
	return controllers.NewSuccessResponse(c,response.FromDomain(followUserr))
	
}