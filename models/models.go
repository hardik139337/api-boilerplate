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
	Id      string `json:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Seats   []Seat `json:"seats" binding:"required"`
}

type Seat struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Number    int    `json:"number"`
	TheaterID string `json:"theater_id"`
}
