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

	_dbDriver "infion-BE/drivers/mysql"

	_middleware "infion-BE/app/middleware"
	_routes "infion-BE/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
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

	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUseCase.NewUseCase(userRepo,timeoutContext)
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

	routesInit := _routes.ControllerList{
		UserController: 		*userCtrl,
		RolesController:		*rolesCtrl,
		ThreadsController:		*threadsCtrl,
		FollowThreadsController:	*followThreadsCtrl,
		LikeThreadsController:		*likeThreadsCtrl,
	}
	routesInit.RouteRegister(e)

	_middleware.LogMiddlewareInit(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}