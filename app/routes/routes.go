package routes

import (
	"infion-BE/controllers/announcements"
	"infion-BE/controllers/comments"
	"infion-BE/controllers/followThreads"
	"infion-BE/controllers/followUsers"
	"infion-BE/controllers/likeComments"
	"infion-BE/controllers/likeReplies"
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
	FollowUsersController 	followUsers.FollowUsersController
	RolesController         roles.RolesController
	ThreadsController       threads.ThreadsController
	FollowThreadsController followThreads.FollowThreadsController
	LikeThreadsController   likeThreads.LikeThreadsController
	CommentsController      comments.CommentsController
	LikeCommentsController  likeComments.LikeCommentsController
	ReportsController       reports.ReportsController
	RepliesController       replies.RepliesController
	LikeRepliesController	likeReplies.LikeRepliesController
	JWTConfig				middleware.JWTConfig
	AnnouncementsController	announcements.AnnouncementsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("user")
	users.POST("/login", cl.UserController.Login)
	users.POST("/create", cl.UserController.CreateNewUser)
	users.GET("/:id", cl.UserController.FindById,middleware.JWTWithConfig(cl.JWTConfig))
	users.PUT("/:id", cl.UserController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	users.GET("/leaderboard/", cl.UserController.GetLeaderboard)

	followUsers := e.Group("followUsers")
	followUsers.POST("/create", cl.FollowUsersController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	followUsers.GET("/:id", cl.FollowUsersController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	followUsers.PUT("/:id", cl.FollowUsersController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	followUsers.DELETE("/:id", cl.FollowUsersController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	roles := e.Group("roles")
	roles.POST("/create", cl.RolesController.Create)
	roles.GET("/:id", cl.RolesController.ReadID)
	roles.PUT("/:id", cl.RolesController.Update)
	roles.DELETE("/:id", cl.RolesController.Delete)

	threads := e.Group("threads")
	threads.POST("/create", cl.ThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	threads.GET("/:id", cl.ThreadsController.ReadID)
	threads.PUT("/:id", cl.ThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	threads.DELETE("/:id", cl.ThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	threads.GET("/list/", cl.ThreadsController.GetThreads)
	threads.GET("/listbyuser/:id", cl.ThreadsController.GetThreadsByUserID)

	followThreads := e.Group("followThreads")
	followThreads.POST("/create", cl.FollowThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.GET("/:id", cl.FollowThreadsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.PUT("/:id", cl.FollowThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.DELETE("/:id", cl.FollowThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	followThreads.GET("/status/", cl.FollowThreadsController.GetStatus,middleware.JWTWithConfig(cl.JWTConfig))

	likeThreads := e.Group("likeThreads")
	likeThreads.POST("/create", cl.LikeThreadsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.GET("/:id", cl.LikeThreadsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.PUT("/:id", cl.LikeThreadsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.DELETE("/:id", cl.LikeThreadsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	likeThreads.GET("/status/", cl.LikeThreadsController.GetStatus,middleware.JWTWithConfig(cl.JWTConfig))

	comments := e.Group("comments")
	comments.POST("/create", cl.CommentsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	comments.GET("/:id", cl.CommentsController.ReadID)
	comments.PUT("/:id", cl.CommentsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	comments.DELETE("/:id", cl.CommentsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	comments.GET("/list", cl.CommentsController.GetComments)
	comments.GET("/listbythread/:id", cl.CommentsController.GetCommentsByThreadID)

	replies := e.Group("replies")
	replies.POST("/create", cl.RepliesController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	replies.GET("/:id", cl.RepliesController.ReadID)
	replies.PUT("/:id", cl.RepliesController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	replies.DELETE("/:id", cl.RepliesController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	replies.GET("/list", cl.RepliesController.GetReplies)
	replies.GET("/listbycomment/:id", cl.RepliesController.GetRepliesByCommentID)

	likeComments := e.Group("likeComments")
	likeComments.POST("/create", cl.LikeCommentsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.GET("/:id", cl.LikeCommentsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.PUT("/:id", cl.LikeCommentsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	likeComments.DELETE("/:id", cl.LikeCommentsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	likeReplies := e.Group("likeReplies")
	likeReplies.POST("/create", cl.LikeRepliesController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	likeReplies.GET("/:id", cl.LikeRepliesController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	likeReplies.PUT("/:id", cl.LikeRepliesController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	likeReplies.DELETE("/:id", cl.LikeRepliesController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

	reports := e.Group("reports")
	reports.POST("/create", cl.ReportsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	reports.GET("/:id", cl.ReportsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	reports.PUT("/:id", cl.ReportsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	reports.DELETE("/:id", cl.ReportsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	reports.GET("/list", cl.ReportsController.GetReports,middleware.JWTWithConfig(cl.JWTConfig))
	reports.GET("/listbyuser/:id", cl.ReportsController.GetReportsByUserID,middleware.JWTWithConfig(cl.JWTConfig))

	announcements := e.Group("announcements")
	announcements.POST("/create", cl.AnnouncementsController.Create,middleware.JWTWithConfig(cl.JWTConfig))
	announcements.GET("/:id", cl.AnnouncementsController.ReadID,middleware.JWTWithConfig(cl.JWTConfig))
	announcements.PUT("/:id", cl.AnnouncementsController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	announcements.DELETE("/:id", cl.AnnouncementsController.Delete,middleware.JWTWithConfig(cl.JWTConfig))
	announcements.GET("/list", cl.AnnouncementsController.GetAnnouncements,middleware.JWTWithConfig(cl.JWTConfig))
	announcements.GET("/listbyuser/:id", cl.AnnouncementsController.GetAnnouncementsByUserID,middleware.JWTWithConfig(cl.JWTConfig))
}
