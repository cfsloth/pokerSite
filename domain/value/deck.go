package value

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"../entity"
)

type Deck struct {
	Id        int64
	CardArray []entity.Card
}

func CreatePokerDeck() Deck {
	fmt.Println("Creating a Poker Deck")
	var deck Deck
	deck.Id = 1
	var array []string
	array = append(array, "Clubs")
	array = append(array, "Hearths")
	array = append(array, "Spaces")
	array = append(array, "Diamonds")
	var id int64 = 1
	for a := 0; a < 4; a++ {
		for i := 1; i < 11; i++ {
			deck.CardArray = append(deck.CardArray, entity.Card{id, strconv.Itoa(i), array[a]})
			id++
		}
		deck.CardArray = append(deck.CardArray, entity.Card{id, "J", array[a]})
		id++
		deck.CardArray = append(deck.CardArray, entity.Card{id, "Q", array[a]})
		id++
		deck.CardArray = append(deck.CardArray, entity.Card{id, "K", array[a]})
		id++
	}
	return deck
}

//The deck should be shuffle 7 times to insure randomness
func RiffleShuffleDeck(deck Deck) Deck {
	var card []entity.Card = deck.CardArray
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)
	randomCounter := 0
	bigPile := 26 + number
	bigPileArray := make([]entity.Card, bigPile)
	smallPileArray := make([]entity.Card, (52 - bigPile))
	copy(smallPileArray, card[bigPile-1:])
	copy(bigPileArray, card[:bigPile-1])
	for i := 0; i < len(smallPileArray); i++ {
		rand.Seed(time.Now().UnixNano())
		slideWindow := rand.Intn(3)
		if len(bigPileArray)-1 <= randomCounter {
			panic("The index excess the size of the array!")
		}
		bigPileArray = append(bigPileArray, entity.Card{})
		copy(bigPileArray[randomCounter+1:], bigPileArray[randomCounter:])
		bigPileArray[randomCounter] = smallPileArray[i]
		randomCounter = randomCounter + slideWindow
	}
	bigPileArray[len(bigPileArray)-1] = smallPileArray[len(smallPileArray)-1]
	deck.CardArray = bigPileArray
	bigPileArray = nil
	smallPileArray = nil
	card = nil
	return deck
}
