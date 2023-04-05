package booking

import (
	"fmt"
	"lcode/models"
	"lcode/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Bookingcontroller struct {
	Repository repository.RepositoryI
}

// @Success      201   {object}  models.Booking
// @Router /bookings [post]
// @Param        booking  body      models.BookingR  true  "Booking JSON"
func (m *Bookingcontroller) AddBooking(c *gin.Context) {
	var Booking models.BookingR

	err := c.ShouldBindJSON(&Booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var tickets []models.Tickets

	err = m.Repository.Raw(&tickets, "select * from tickets where seat_id IN (?) and show_id = ?",
		Booking.SeatIdS, Booking.ShowId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	fmt.Printf("tickets: %v\n", tickets)
	var booking = &models.Booking{UserId: Booking.UserId, Id: uuid.New().String()}
	m.Repository.Create(booking)
	var ticketBookingArr []models.TicketsBooking
	for _, v := range tickets {
		ticketBookingArr = append(ticketBookingArr, models.TicketsBooking{
			Id:        uuid.New().String(),
			BookingId: booking.Id,
			TicketId:  v.Id,
		})
	}
	m.Repository.Create(&ticketBookingArr)

	c.JSON(http.StatusCreated, Booking)
}

// @Router /Bookings [get]
// @Success      200   {object}  []models.Booking
func (m *Bookingcontroller) GetBookings(c *gin.Context) {
	var Bookings []models.Booking
	err := m.Repository.QueryAll(&Bookings, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Bookings)
}

// @Router /Bookings/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Booking
func (m *Bookingcontroller) GetByIdBookings(c *gin.Context) {
	var Booking models.Booking
	id := c.Param("id")
	Booking.Id = id
	err := m.Repository.Query(&Booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Booking)
}

// @Router /Bookings/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Booking
func (m *Bookingcontroller) UpdateBooking(c *gin.Context) {
	var Booking models.Booking
	id := c.Param("id")
	Booking.Id = id
	err := c.ShouldBindJSON(&Booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var BookingModel models.Booking = models.Booking(Booking)

	err = m.Repository.Update(&BookingModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Booking)
}

// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Booking
// @Router /Bookings/{id} [delete]
func (m *Bookingcontroller) DeleteBooking(c *gin.Context) {
	var Booking models.Booking
	id := c.Param("id")
	Booking.Id = id
	err := c.ShouldBindJSON(&Booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	BookingModel := models.Booking(Booking)
	err = m.Repository.Delete(&BookingModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Booking)
}
