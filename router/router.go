package router

import (
	"lcode/constants"
	"lcode/controller"
	"lcode/models"
	"lcode/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db.AutoMigrate(&models.Movie{})
	if err != nil {
		panic(err)
	}
	repository := repository.Repository{Db: db}
	moviec := controller.Moviecontroller{
		Repository: &repository,
	}
	r.POST(constants.Movies, moviec.AddMovie)
	r.GET(constants.Movies, moviec.GetMovies)
	r.PATCH(constants.Movies, moviec.UpdateMovie)
	r.DELETE(constants.Movies, moviec.DeleteMovie)

	return r
}
