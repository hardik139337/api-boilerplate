package router

import (
	"lcode/constants"

	controller "lcode/controller/movie"
	theater "lcode/controller/theater"
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

	addRoute(r)
	return r
}

func addRoute(r *gin.Engine) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(models.Movie), new(models.Theater), new(models.Seat))

	repository := repository.Repository{Db: db}
	moviec := controller.Moviecontroller{
		Repository: &repository,
	}

	routeM := r.Group(constants.Movies)
	{
		routeM.POST("", moviec.AddMovie)
		routeM.GET("", moviec.GetMovies)
		routeM.GET(constants.ID, moviec.GetByIdMovies)
		routeM.PATCH(constants.ID, moviec.UpdateMovie)
		routeM.DELETE(constants.ID, moviec.DeleteMovie)
	}

	theaterc := theater.Theatercontroller{
		Repository: &repository,
	}
	routeT := r.Group(constants.Theaters)
	{
		routeT.POST("", theaterc.AddTheater)
		routeT.GET("", theaterc.GetTheaters)
		routeT.GET(constants.ID, theaterc.GetByIdTheaters)
		routeT.PATCH(constants.ID, theaterc.UpdateTheater)
		routeT.DELETE(constants.ID, theaterc.DeleteTheater)
	}
}
