package controller

import (
	"lcode/models"
	"lcode/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Moviecontroller struct {
	Repository repository.RepositoryI
}

// @Success      201   {object}  models.Movie
// @Router /movies [post]
// @Param        Movie  body      models.Movie  true  "Movie JSON"
func (m *Moviecontroller) AddMovie(c *gin.Context) {
	var movie models.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movie.Id = uuid.New().String()

	err = m.Repository.Create(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// @Router /movies [get]
// @Success      200   {object}  []models.Movie
func (m *Moviecontroller) GetMovies(c *gin.Context) {
	var movies []models.Movie
	err := m.Repository.QueryAll(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// @Router /movies/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
func (m *Moviecontroller) GetByIdMovies(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")
	movie.Id = id
	err := m.Repository.Query(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// @Router /movies/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
func (m *Moviecontroller) UpdateMovie(c *gin.Context) {
	var movie models.MovieUpdate
	id := c.Param("id")
	movie.Id = id
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var movieModel models.Movie = models.Movie(movie)

	err = m.Repository.Update(&movieModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
// @Router /movies/{id} [delete]
func (m *Moviecontroller) DeleteMovie(c *gin.Context) {
	var movie models.MovieUpdate
	id := c.Param("id")
	movie.Id = id
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movieModel := models.Movie(movie)
	err = m.Repository.Delete(&movieModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
