package entity

type Hand struct{
    FirstCard Card 'json:"firstCard"'
    SecondCard Card 'json:"secondCard"'
}
