package router

import (
	"lcode/constants"

	"lcode/controller/booking"
	controller "lcode/controller/movie"
	"lcode/controller/show"
	theater "lcode/controller/theater"
	"lcode/models"
	"lcode/repository"

	"github.com/gin-gonic/gin"

	_ "lcode/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	addRoute(r)
	return r
}

func addRoute(r *gin.Engine) {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(models.Movie), new(models.Theater), new(models.Seat), new(models.Show),
		new(models.Tickets), new(models.Booking), new(models.TicketsBooking))

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

	theaterS := show.Showcontroller{
		Repository: &repository,
	}
	routeS := r.Group(constants.SHOWS)
	{
		routeS.POST("", theaterS.AddShow)
		routeS.GET("", theaterS.Getshows)
		routeS.GET(constants.ID, theaterS.GetByIdshows)
		routeS.PATCH(constants.ID, theaterS.Updateshow)
		routeS.DELETE(constants.ID, theaterS.Deleteshow)
	}

	bookingC := booking.Bookingcontroller{
		Repository: &repository,
	}
	routeB := r.Group(constants.BOOKING)
	{
		routeB.POST("", bookingC.AddBooking)
		routeB.GET("", bookingC.GetBookings)
		routeB.GET(constants.ID, bookingC.GetByIdBookings)
		routeB.PATCH(constants.ID, bookingC.UpdateBooking)
		routeB.DELETE(constants.ID, bookingC.DeleteBooking)
	}
}
