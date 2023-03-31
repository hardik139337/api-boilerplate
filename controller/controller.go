package controller

import (
	"fmt"
	"lcode/models"
	"lcode/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Moviecontroller struct {
	Repository repository.RepositoryI
}

func (m *Moviecontroller) AddMovie(c *gin.Context) {
	var movie models.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movie.Id = uuid.New().String()
	fmt.Println(movie, "movie")

	err = m.Repository.Create(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	fmt.Println(movie, "movie")
	c.JSON(http.StatusCreated, movie)
}

func (m *Moviecontroller) GetMovies(c *gin.Context) {
	var movies []models.Movie
	err := m.Repository.QueryAll(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (m *Moviecontroller) UpdateMovie(c *gin.Context) {
	var movie models.MovieUpdate
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	err = m.Repository.Update(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *Moviecontroller) DeleteMovie(c *gin.Context) {
	var movie models.MovieUpdate
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	err = m.Repository.Delete(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
