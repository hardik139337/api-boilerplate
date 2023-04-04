package models

type Movie struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Language string `json:"language" binding:"required"`
	Length   int64  `json:"length" binding:"required"`
}

type MovieUpdate struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Length   int64  `json:"length"`
}

type Error struct {
	Message string `json:"message"`
}

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Theater struct {
	Id      string `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" `
	Address string `json:"address" `
	Seats   []Seat `json:"seats" `
}

type Seat struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Number    int    `json:"number"`
	TheaterID string `json:"theater_id"`
}

type Show struct {
	Id        string `json:"id" gorm:"primaryKey"`
	MovieId   string `json:"MovieId" binding:"required"`
	TheaterId string `json:"TheaterId" binding:"required"`
	StartTime int64  `json:"StartTime" binding:"required"`
}

type Tickets struct {
	Id     string `json:"id" gorm:"primaryKey"`
	Price  int    `json:"Price"`
	ShowId string `json:"ShowId"`
	SeatId string `json:"SeatId"`
}

type Booking struct {
	Id     string `json:"id" gorm:"primaryKey"`
	UserId string
}

type TicketsBooking struct {
	Id        string `json:"id" gorm:"primaryKey"`
	BookingId string
	TicketId  string
}

type BookingR struct {
	UserId  string   `json:"UserId" binding:"required"`
	ShowId  string   `json:"ShowId" binding:"required"`
	SeatIdS []string `json:"SeatIdS" binding:"required"`
}
