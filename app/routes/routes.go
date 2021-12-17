package routes

import (
	"infion-BE/controllers/threads"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	ThreadsController     threads.ThreadsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create)
	threads.GET("/:id", cl.ThreadsController.ReadID)
	// threads.PUT("/update", cl.ThreadsController.Update)
}