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
		&_rolesRepo.Roles{},
		&_threadsRepo.Threads{},
		&_followThreadsRepo.FollowThreads{},
		&_likeThreadsRepo.LikeThreads{},
		&_commentsRepo.Comments{},
		&_likeCommentsRepo.LikeComments{},
		&_reportsRepo.Reports{},
		&_repliesRepo.Replies{},
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

	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUseCase.NewUseCase(userRepo, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	rolesRepo := _rolesRepo.NewRolesRepository(db)
	rolesUsecase := _rolesUsecase.NewRolesUsecase(rolesRepo, timeoutContext)
	rolesCtrl := _rolesController.NewRolesController(rolesUsecase)

	threadsRepo := _threadsRepo.NewThreadsRepository(db)
	threadsUsecase := _threadsUsecase.NewThreadsUsecase(threadsRepo, timeoutContext)
	threadsCtrl := _threadsController.NewThreadsController(threadsUsecase)

	followThreadsRepo := _followThreadsRepo.NewFollowThreadsRepository(db)
	followThreadsUsecase := _followThreadsUsecase.NewFollowThreadsUsecase(followThreadsRepo, timeoutContext)
	followThreadsCtrl := _followThreadsController.NewFollowThreadsController(followThreadsUsecase)

	likeThreadsRepo := _likeThreadsRepo.NewLikeThreadsRepository(db)
	likeThreadsUsecase := _likeThreadsUsecase.NewLikeThreadsUsecase(likeThreadsRepo, timeoutContext)
	likeThreadsCtrl := _likeThreadsController.NewLikeThreadsController(likeThreadsUsecase)

	repliesRepo := _repliesRepo.NewRepliesRepository(db)
	repliesUsecase := _repliesUsecase.NewRepliesUsecase(repliesRepo, timeoutContext)
	repliesCtrl := _repliesController.NewRepliesController(repliesUsecase)

	commentsRepo := _commentsRepo.NewCommentsRepository(db)
	commentsUsecase := _commentsUsecase.NewCommentsUsecase(commentsRepo, timeoutContext, repliesRepo)
	commentsCtrl := _commentsController.NewCommentsController(commentsUsecase)

	likeCommentsRepo := _likeCommentsRepo.NewLikeCommentsRepository(db)
	likeCommentsUsecase := _likeCommentsUsecase.NewLikeCommentsUsecase(likeCommentsRepo, timeoutContext)
	likeCommentsCtrl := _likeCommentsController.NewLikeCommentsController(likeCommentsUsecase)

	reportsRepo := _reportsRepo.NewReportsRepository(db)
	reportsUsecase := _reportsUsecase.NewReportsUsecase(reportsRepo, timeoutContext)
	reportsCtrl := _reportsController.NewReportsController(reportsUsecase)

	routesInit := _routes.ControllerList{
		UserController:          *userCtrl,
		RolesController:         *rolesCtrl,
		ThreadsController:       *threadsCtrl,
		FollowThreadsController: *followThreadsCtrl,
		LikeThreadsController:   *likeThreadsCtrl,
		CommentsController:      *commentsCtrl,
		LikeCommentsController:  *likeCommentsCtrl,
		ReportsController:       *reportsCtrl,
		RepliesController:       *repliesCtrl,
	}
	routesInit.RouteRegister(e)

	_middleware.LogMiddlewareInit(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
