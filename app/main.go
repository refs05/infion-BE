package main

import (
	_threadsUsecase "infion-BE/businesses/threads"
	_threadsController "infion-BE/controllers/threads"
	_threadsRepo "infion-BE/drivers/databases/threads"

	_rolesUsecase "infion-BE/businesses/roles"
	_rolesController "infion-BE/controllers/roles"
	_rolesRepo "infion-BE/drivers/databases/roles"

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
		&_threadsRepo.Threads{},
		&_rolesRepo.Roles{},
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

	threadsRepo := _threadsRepo.NewThreadsRepository(db)
	threadsUsecase := _threadsUsecase.NewThreadsUsecase(threadsRepo, timeoutContext)
	threadsCtrl := _threadsController.NewThreadsController(threadsUsecase)

	rolesRepo := _rolesRepo.NewRolesRepository(db)
	rolesUsecase := _rolesUsecase.NewRolesUsecase(rolesRepo, timeoutContext)
	rolesCtrl := _rolesController.NewRolesController(rolesUsecase)


	routesInit := _routes.ControllerList{
		ThreadsController:		*threadsCtrl,
		RolesController:		*rolesCtrl,
	}
	routesInit.RouteRegister(e)

	_middleware.LogMiddlewareInit(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}