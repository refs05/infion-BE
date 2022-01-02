package routes

import (
	"infion-BE/controllers/comments"
	"infion-BE/controllers/followThreads"
	"infion-BE/controllers/likeThreads"
	"infion-BE/controllers/reports"
	"infion-BE/controllers/roles"
	"infion-BE/controllers/threads"

	userController "infion-BE/controllers/users"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController				userController.UserController
	RolesController				roles.RolesController
	ThreadsController			threads.ThreadsController
	FollowThreadsController		followThreads.FollowThreadsController
	LikeThreadsController		likeThreads.LikeThreadsController
	CommentsController			comments.CommentsController
	ReportsController			reports.ReportsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("user")
	users.POST("/login",cl.UserController.Login)
	users.POST("/create",cl.UserController.CreateNewUser)
	users.GET("/:id",cl.UserController.FindById)

	roles := e.Group("roles")
	roles.POST("/create", cl.RolesController.Create)
	roles.GET("/:id", cl.RolesController.ReadID)
	roles.PUT("/:id", cl.RolesController.Update)
	roles.DELETE("/:id", cl.RolesController.Delete)

	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create)
	threads.GET("/:id", cl.ThreadsController.ReadID)
	threads.PUT("/:id", cl.ThreadsController.Update)
	threads.DELETE("/:id", cl.ThreadsController.Delete)
	threads.GET("/list", cl.ThreadsController.GetThreads)

	followThreads := e.Group("followThreads")
	followThreads.POST("/create", cl.FollowThreadsController.Create)
	followThreads.GET("/:id", cl.FollowThreadsController.ReadID)
	followThreads.PUT("/:id", cl.FollowThreadsController.Update)
	followThreads.DELETE("/:id", cl.FollowThreadsController.Delete)

	likeThreads := e.Group("likeThreads")
	likeThreads.POST("/create", cl.LikeThreadsController.Create)
	likeThreads.GET("/:id", cl.LikeThreadsController.ReadID)
	likeThreads.PUT("/:id", cl.LikeThreadsController.Update)
	likeThreads.DELETE("/:id", cl.LikeThreadsController.Delete)

	comments := e.Group("comments")
	comments.POST("/create", cl.CommentsController.Create)
	comments.GET("/:id", cl.CommentsController.ReadID)
	comments.PUT("/:id", cl.CommentsController.Update)
	comments.DELETE("/:id", cl.CommentsController.Delete)
	comments.GET("/list", cl.CommentsController.GetComments)

	reports := e.Group("reports")
	reports.POST("/create", cl.ReportsController.Create)
	reports.GET("/:id", cl.ReportsController.ReadID)
	reports.PUT("/:id", cl.ReportsController.Update)
	reports.DELETE("/:id", cl.ReportsController.Delete)
	reports.GET("/list", cl.ReportsController.GetReports)
}