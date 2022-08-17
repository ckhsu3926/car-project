package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	"gorm.io/gorm"

	_ "car-record/docs"

	"car-record/config"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_middleware "car-record/middleware"

	_userDeliveryHttp "car-record/domain/User/delivery/http"
	_userRepositoryMysql "car-record/domain/User/repository/mysql"
	_userUsecase "car-record/domain/User/usecase"
	_vehicleDeliveryHttp "car-record/domain/Vehicle/delivery/http"
	_vehicleRepositoryMysql "car-record/domain/Vehicle/repository/mysql"
	_vehicleUsecase "car-record/domain/Vehicle/usecase"
)

var mysqlConnection *gorm.DB
var timeContext = time.Duration(2) * time.Second

func init() {
	// env
	if err := config.InitialEnvConfiguration(); err != nil {
		panic(err)
	}

	// database
	mysqlConnection = config.NewMysqlConnection(config.EnvConfig.ConnectionStrings.Mysql)
}

// @Title  Car-Record
// @Description
// @Version  0.1
// @Host     localhost:8081
func main() {
	// gin
	if config.EnvConfig.Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// middleware
	corsMiddleware := _middleware.NewCORSMiddleware()
	r.Use(corsMiddleware.CORS())

	// api
	apiRouter := r.Group("api")

	userRepo := _userRepositoryMysql.NewMysqlUserRepository(mysqlConnection)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeContext)
	_userDeliveryHttp.NewUserHttpHandler(apiRouter, userUsecase)

	// need authorize api
	authorizeMiddleware := _middleware.NewAuthorizeMiddleware(userUsecase)
	authorizedApiRouter := apiRouter.Group("", authorizeMiddleware.Authorize())

	vehicleRepo := _vehicleRepositoryMysql.NewMysqlVehicleRepository(mysqlConnection)
	vehicleUsecase := _vehicleUsecase.NewVehicleUsecase(vehicleRepo, timeContext)
	_vehicleDeliveryHttp.NewVehicleHttpHandler(authorizedApiRouter.Group("vehicle"), vehicleUsecase)

	// gin swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// setup frontend dist spa
	r.Use(spa.Middleware("/", config.EnvConfig.FrontEndDir))

	if err := r.Run("localhost:" + config.EnvConfig.Port); err != nil {
		panic(err)
	}
}
