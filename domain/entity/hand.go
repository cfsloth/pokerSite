package entity

type Hand struct {
	ID         int    `json:"Id"`
	PeoplesID  string `json:"peopleID"`
	FirstCard  Card   `json:"firstCard"`
	SecondCard Card   `json:"secondCard"`
}
