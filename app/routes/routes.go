package routes

import (
	"infion-BE/controllers/roles"
	"infion-BE/controllers/threads"

	echo "github.com/labstack/echo/v4"
	userController "infion-BE/controllers/users"
)

type ControllerList struct {
	ThreadsController    threads.ThreadsController
	RolesController     roles.RolesController
	UserController 		userController.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create)
	threads.GET("/:id", cl.ThreadsController.ReadID)
	threads.PUT("/:id", cl.ThreadsController.Update)
	threads.DELETE("/:id", cl.ThreadsController.Delete)
	threads.GET("/list", cl.ThreadsController.GetThreads)

	users := e.Group("user")
	users.POST("/login",cl.UserController.Login)
	users.POST("/create",cl.UserController.CreateNewUser)
	users.GET("/:id",cl.UserController.FindById)

	roles := e.Group("roles")
	roles.POST("/create", cl.RolesController.Create)
	roles.GET("/:id", cl.RolesController.ReadID)
	roles.PUT("/:id", cl.RolesController.Update)
	roles.DELETE("/:id", cl.RolesController.Delete)
}