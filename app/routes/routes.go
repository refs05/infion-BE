package routes

import (
	"infion-BE/controllers/roles"
	"infion-BE/controllers/threads"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	ThreadsController     threads.ThreadsController
	RolesController     roles.RolesController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create)
	threads.GET("/:id", cl.ThreadsController.ReadID)
	threads.PUT("/:id", cl.ThreadsController.Update)
	threads.DELETE("/:id", cl.ThreadsController.Delete)

	roles := e.Group("roles")
	roles.POST("/create", cl.RolesController.Create)
	roles.GET("/:id", cl.RolesController.ReadID)
	roles.PUT("/:id", cl.RolesController.Update)
	roles.DELETE("/:id", cl.RolesController.Delete)
}