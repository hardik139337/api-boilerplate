package theater

import (
	"lcode/models"
	"lcode/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Theatercontroller struct {
	Repository repository.RepositoryI
}

// @Success      201   {object}  models.Theater
// @Router /theaters [post]
// @Param        Theater  body      models.Theater  true  "Theater JSON"
func (m *Theatercontroller) AddTheater(c *gin.Context) {
	var Theater models.Theater

	err := c.ShouldBindJSON(&Theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	Theater.Id = uuid.New().String()

	err = m.Repository.Create(&Theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Theater)
}

// @Router /theaters [get]
// @Success      200   {object}  []models.Theater
func (m *Theatercontroller) GetTheaters(c *gin.Context) {
	var TheatersFilter models.Theater
	err := c.BindQuery(&TheatersFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var Theaters []models.Theater
	err = m.Repository.QueryAll(&Theaters, TheatersFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theaters)
}

// @Router /theaters/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
func (m *Theatercontroller) GetByIdTheaters(c *gin.Context) {
	var Theater models.Theater
	id := c.Param("id")
	Theater.Id = id
	err := m.Repository.Query(&Theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theater)
}

// @Router /theaters/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
// @Param        Theater  body      models.Theater  true  "Theater JSON"
func (m *Theatercontroller) UpdateTheater(c *gin.Context) {
	var theater models.Theater
	err := c.ShouldBindJSON(&theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	id := c.Param("id")
	theater.Id = id

	var TheaterModel models.Theater = models.Theater(theater)

	err = m.Repository.Update(&TheaterModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, theater)
}

// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
// @Router /theaters/{id} [delete]
func (m *Theatercontroller) DeleteTheater(c *gin.Context) {
	var Theater models.Theater
	id := c.Param("id")
	Theater.Id = id
	TheaterModel := models.Theater(Theater)
	err := m.Repository.Delete(&TheaterModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theater)
}
