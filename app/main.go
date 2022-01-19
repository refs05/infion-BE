package main

import (
	_userUseCase "infion-BE/businesses/users"
	_userController "infion-BE/controllers/users"
	_userRepo "infion-BE/drivers/databases/users"

	_rolesUsecase "infion-BE/businesses/roles"
	_rolesController "infion-BE/controllers/roles"
	_rolesRepo "infion-BE/drivers/databases/roles"

	_threadsUsecase "infion-BE/businesses/threads"
	_threadsController "infion-BE/controllers/threads"
	_threadsRepo "infion-BE/drivers/databases/threads"

	_followThreadsUsecase "infion-BE/businesses/followThreads"
	_followThreadsController "infion-BE/controllers/followThreads"
	_followThreadsRepo "infion-BE/drivers/databases/followThreads"

	_followUsersUsecase "infion-BE/businesses/followUsers"
	_followUsersController "infion-BE/controllers/followUsers"
	_followUsersRepo "infion-BE/drivers/databases/followUsers"

	_likeThreadsUsecase "infion-BE/businesses/likeThreads"
	_likeThreadsController "infion-BE/controllers/likeThreads"
	_likeThreadsRepo "infion-BE/drivers/databases/likeThreads"

	_commentsUsecase "infion-BE/businesses/comments"
	_commentsController "infion-BE/controllers/comments"
	_commentsRepo "infion-BE/drivers/databases/comments"

	_repliesUsecase "infion-BE/businesses/replies"
	_repliesController "infion-BE/controllers/replies"
	_repliesRepo "infion-BE/drivers/databases/replies"

	_likeCommentsUsecase "infion-BE/businesses/likeComments"
	_likeCommentsController "infion-BE/controllers/likeComments"
	_likeCommentsRepo "infion-BE/drivers/databases/likeComments"

	_likeRepliesUsecase "infion-BE/businesses/likeReplies"
	_likeRepliesController "infion-BE/controllers/likeReplies"
	_likeRepliesRepo "infion-BE/drivers/databases/likeReplies"

	_reportsUsecase "infion-BE/businesses/reports"
	_reportsController "infion-BE/controllers/reports"
	_reportsRepo "infion-BE/drivers/databases/reports"

	_dbDriver "infion-BE/drivers/mysql"

	_middleware "infion-BE/app/middleware"
	_routes "infion-BE/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {

	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.User{},
		&_followUsersRepo.FollowUsers{},
		&_rolesRepo.Roles{},
		&_threadsRepo.Threads{},
		&_followThreadsRepo.FollowThreads{},
		&_likeThreadsRepo.LikeThreads{},
		&_commentsRepo.Comments{},
		&_likeCommentsRepo.LikeComments{},
		&_reportsRepo.Reports{},
		&_repliesRepo.Replies{},
		&_likeRepliesRepo.LikeReplies{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	e.Use(middleware.CORS())

	rolesRepo := _rolesRepo.NewRolesRepository(db)
	rolesUsecase := _rolesUsecase.NewRolesUsecase(rolesRepo, timeoutContext)
	rolesCtrl := _rolesController.NewRolesController(rolesUsecase)

	likeRepliesRepo := _likeRepliesRepo.NewLikeRepliesRepository(db)
	likeRepliesUsecase := _likeRepliesUsecase.NewLikeRepliesUsecase(likeRepliesRepo, timeoutContext)
	likeRepliesCtrl := _likeRepliesController.NewLikeRepliesController(likeRepliesUsecase)

	repliesRepo := _repliesRepo.NewRepliesRepository(db)
	repliesUsecase := _repliesUsecase.NewRepliesUsecase(repliesRepo, timeoutContext, likeRepliesRepo)
	repliesCtrl := _repliesController.NewRepliesController(repliesUsecase)

	likeCommentsRepo := _likeCommentsRepo.NewLikeCommentsRepository(db)
	likeCommentsUsecase := _likeCommentsUsecase.NewLikeCommentsUsecase(likeCommentsRepo, timeoutContext)
	likeCommentsCtrl := _likeCommentsController.NewLikeCommentsController(likeCommentsUsecase)

	commentsRepo := _commentsRepo.NewCommentsRepository(db)
	commentsUsecase := _commentsUsecase.NewCommentsUsecase(commentsRepo, timeoutContext, repliesRepo, likeCommentsRepo)
	commentsCtrl := _commentsController.NewCommentsController(commentsUsecase)

	followThreadsRepo := _followThreadsRepo.NewFollowThreadsRepository(db)
	followThreadsUsecase := _followThreadsUsecase.NewFollowThreadsUsecase(followThreadsRepo, timeoutContext)
	followThreadsCtrl := _followThreadsController.NewFollowThreadsController(followThreadsUsecase)

	likeThreadsRepo := _likeThreadsRepo.NewLikeThreadsRepository(db)
	likeThreadsUsecase := _likeThreadsUsecase.NewLikeThreadsUsecase(likeThreadsRepo, timeoutContext)
	likeThreadsCtrl := _likeThreadsController.NewLikeThreadsController(likeThreadsUsecase)

	threadsRepo := _threadsRepo.NewThreadsRepository(db)
	threadsUsecase := _threadsUsecase.NewThreadsUsecase(threadsRepo, timeoutContext, likeThreadsRepo, commentsRepo, followThreadsRepo)
	threadsCtrl := _threadsController.NewThreadsController(threadsUsecase)

	followUsersRepo := _followUsersRepo.NewFollowUsersRepository(db)
	followUsersUsecase := _followUsersUsecase.NewFollowUsersUsecase(followUsersRepo, timeoutContext)
	followUsersCtrl := _followUsersController.NewFollowUsersController(followUsersUsecase)

	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUseCase.NewUseCase(userRepo, timeoutContext, commentsRepo)
	userCtrl := _userController.NewUserController(userUsecase)

	reportsRepo := _reportsRepo.NewReportsRepository(db)
	reportsUsecase := _reportsUsecase.NewReportsUsecase(reportsRepo, timeoutContext)
	reportsCtrl := _reportsController.NewReportsController(reportsUsecase)

	routesInit := _routes.ControllerList{
		UserController:				*userCtrl,
		FollowUsersController:		*followUsersCtrl,
		RolesController:			*rolesCtrl,
		ThreadsController:			*threadsCtrl,
		FollowThreadsController:	*followThreadsCtrl,
		LikeThreadsController:		*likeThreadsCtrl,
		CommentsController:			*commentsCtrl,
		LikeCommentsController:		*likeCommentsCtrl,
		ReportsController:			*reportsCtrl,
		RepliesController:			*repliesCtrl,
		LikeRepliesController:		*likeRepliesCtrl,
	}
	routesInit.RouteRegister(e)

	_middleware.LogMiddlewareInit(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
