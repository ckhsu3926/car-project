package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	"gorm.io/gorm"

	_ "car-record/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_middleware "car-record/middleware"

	"car-record/config"
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
	apiRouter.GET("test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// gin swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// setup frontend dist spa
	r.Use(spa.Middleware("/", config.EnvConfig.FrontEndDir))

	if err := r.Run("localhost:" + config.EnvConfig.Port); err != nil {
		panic(err)
	}
}
