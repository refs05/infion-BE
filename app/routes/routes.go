package routes

import (
	userController "infion-BE/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController

}

func (controller RouteControllerList)RouteRegister(c *echo.Echo){
	users := c.Group("/user")
	users.POST("/login",controller.UserController.Login)
	users.POST("/create",controller.UserController.CreateNewUser)
}