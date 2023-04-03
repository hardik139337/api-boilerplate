package show

import (
	"fmt"
	"lcode/models"
	"lcode/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Showcontroller struct {
	Repository repository.RepositoryI
}

// @Success      201   {object}  models.Show
// @Router /shows [post]
// @Param        show  body      models.Show  true  "show JSON"
func (m *Showcontroller) AddShow(c *gin.Context) {
	var show models.Show

	err := c.ShouldBindJSON(&show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	show.Id = uuid.New().String()
	showTheater := models.Theater{Id: show.TheaterId}
	err = m.Repository.Query(&showTheater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var showTickets []models.Tickets
	for _, v := range showTheater.Seats {
		showTickets = append(showTickets, models.Tickets{
			SeatId: v.ID,
			Price:  0,
			ShowId: show.Id,
			Id:     uuid.New().String(),
		})
	}
	err = m.Repository.Create(&showTickets)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	err = m.Repository.Create(&show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, show)
}

// @Router /shows [get]
// @Success      200   {object}  []models.Show
func (m *Showcontroller) Getshows(c *gin.Context) {
	var shows []models.Show
	err := m.Repository.QueryAll(&shows, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, shows)
}

// @Router /shows/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
func (m *Showcontroller) GetByIdshows(c *gin.Context) {
	name := c.Query("name") // shortcut for c.Request.URL.Query().Get("lastname")
	fmt.Println(name)
	var show models.Show
	id := c.Param("id")

	show.Id = id
	err := m.Repository.Query(&show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, show)
}

// @Router /shows/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
// @Param        show  body      models.Show  true  "show JSON"
func (m *Showcontroller) Updateshow(c *gin.Context) {
	var show models.Show
	err := c.ShouldBindJSON(&show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	id := c.Param("id")
	show.Id = id

	var showModel models.Show = models.Show(show)

	err = m.Repository.Update(&showModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, show)
}

// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
// @Router /shows/{id} [delete]
func (m *Showcontroller) Deleteshow(c *gin.Context) {
	var show models.Show
	id := c.Param("id")
	show.Id = id
	showModel := models.Show(show)
	err := m.Repository.Delete(&showModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, show)
}
