package routes

import (
	"infion-BE/controllers/comments"
	"infion-BE/controllers/followThreads"
	"infion-BE/controllers/likeComments"
	"infion-BE/controllers/likeThreads"
	"infion-BE/controllers/replies"
	"infion-BE/controllers/reports"
	"infion-BE/controllers/roles"
	"infion-BE/controllers/threads"

	userController "infion-BE/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController          userController.UserController
	RolesController         roles.RolesController
	ThreadsController       threads.ThreadsController
	FollowThreadsController followThreads.FollowThreadsController
	LikeThreadsController   likeThreads.LikeThreadsController
	CommentsController      comments.CommentsController
	LikeCommentsController  likeComments.LikeCommentsController
	ReportsController       reports.ReportsController
	RepliesController       replies.RepliesController
	JWTConfig				middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("user")
	users.POST("/login", cl.UserController.Login)
	users.POST("/create", cl.UserController.CreateNewUser)
	users.GET("/:id", cl.UserController.FindById,middleware.JWTWithConfig(cl.JWTConfig))

	roles := e.Group("roles")
	roles.POST("/create", cl.RolesController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	roles.GET("/:id", cl.RolesController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	roles.PUT("/:id", cl.RolesController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	roles.DELETE("/:id", cl.RolesController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	threads.GET("/:id", cl.ThreadsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	threads.PUT("/:id", cl.ThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	threads.DELETE("/:id", cl.ThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	threads.GET("/list/", cl.ThreadsController.GetThreads,middleware.JWTWithConfig(cl.JWTConfig))

	followThreads := e.Group("followThreads")
	followThreads.POST("/create", cl.FollowThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.GET("/:id", cl.FollowThreadsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.PUT("/:id", cl.FollowThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.DELETE("/:id", cl.FollowThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	likeThreads := e.Group("likeThreads")
	likeThreads.POST("/create", cl.LikeThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.GET("/:id", cl.LikeThreadsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.PUT("/:id", cl.LikeThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.DELETE("/:id", cl.LikeThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	comments := e.Group("comments")
	comments.POST("/create", cl.CommentsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	comments.GET("/:id", cl.CommentsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	comments.PUT("/:id", cl.CommentsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	comments.DELETE("/:id", cl.CommentsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	comments.GET("/list", cl.CommentsController.GetComments,middleware.JWTWithConfig(cl.JWTConfig))
	comments.GET("/listbythread/:id", cl.CommentsController.GetCommentsByThreadID,middleware.JWTWithConfig(cl.JWTConfig))

	replies := e.Group("replies")
	replies.POST("/create", cl.RepliesController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	replies.GET("/:id", cl.RepliesController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	replies.PUT("/:id", cl.RepliesController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	replies.DELETE("/:id", cl.RepliesController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	replies.GET("/list", cl.RepliesController.GetReplies,middleware.JWTWithConfig(cl.JWTConfig))
	replies.GET("/listbycomment/:id", cl.RepliesController.GetRepliesByCommentID,middleware.JWTWithConfig(cl.JWTConfig))

	likeComments := e.Group("likeComments")
	likeComments.POST("/create", cl.LikeCommentsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.GET("/:id", cl.LikeCommentsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.PUT("/:id", cl.LikeCommentsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.DELETE("/:id", cl.LikeCommentsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	reports := e.Group("reports")
	reports.POST("/create", cl.ReportsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	reports.GET("/:id", cl.ReportsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	reports.PUT("/:id", cl.ReportsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	reports.DELETE("/:id", cl.ReportsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	reports.GET("/list", cl.ReportsController.GetReports,middleware.JWTWithConfig(cl.JWTConfig))
}
