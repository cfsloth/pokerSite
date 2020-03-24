package value

import (
  "../entity"
  "strconv"
  "fmt"
)

type Deck struct{
  Id int64 
  CardArray []entity.Card
}

func CreatePokerDeck() Deck{
  fmt.Println("Creating a Poker Deck")
  var deck Deck
  deck.Id = 1
  var array []string
  array = append(array,"Clubs")
  array = append(array,"Hearths")
  array = append(array,"Spaces")
  array = append(array,"Diamonds")
  var id int64 = 1
  for a:=0; a<4; a++ {
      for i:=1; i < 11; i++ {
        deck.CardArray = append(deck.CardArray,entity.Card{id,strconv.Itoa(i),array[a]})
        id++
      }
      deck.CardArray = append(deck.CardArray,entity.Card{id,"J",array[a]})
      id++
      deck.CardArray = append(deck.CardArray,entity.Card{id,"Q",array[a]})
      id++
      deck.CardArray= append(deck.CardArray,entity.Card{id,"K",array[a]})
      id++
  }
  return deck
}
