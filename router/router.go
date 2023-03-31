package router

import (
	"lcode/constants"
	"lcode/controller"
	"lcode/models"
	"lcode/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	_ "lcode/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	addMovieRoute(r)
	return r
}

func addMovieRoute(r *gin.Engine) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Movie{})

	repository := repository.Repository{Db: db}
	moviec := controller.Moviecontroller{
		Repository: &repository,
	}

	routeM := r.Group(constants.Movies)

	routeM.POST("", moviec.AddMovie)
	routeM.GET("", moviec.GetMovies)

	routeM.GET("/:id", moviec.GetByIdMovies)
	routeM.PATCH("/:id", moviec.UpdateMovie)
	routeM.DELETE("/:id", moviec.DeleteMovie)
}
