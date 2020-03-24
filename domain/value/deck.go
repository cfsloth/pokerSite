package value

import (
	"fmt"
	"math/rand"
	"strconv"

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
func RiffleShuffleDeck(deck Deck) {
	var card []entity.Card = deck.CardArray
	number := rand.Intn(10)
	randomCounter := 0
	bigPile := 26 + number
	smallPileArray := card[bigPile-1:]
	bigPileArray := card[:bigPile-1]
	size := len(smallPileArray) + len(bigPileArray)
	fmt.Println(size)
	for i := 0; i < (26-number)+1; i++ {
		slideWindow := rand.Intn(3)

		if len(bigPileArray)-1 <= randomCounter {
			fmt.Println("Special case for when going to the end of the splice")
		}
		bigPileArray = append(bigPileArray, entity.Card{})
		a := randomCounter + 1
		copy(bigPileArray[a:], bigPileArray[randomCounter:])
		bigPileArray[randomCounter] = smallPileArray[i]

		randomCounter += slideWindow

	}
	fmt.Println(bigPileArray)
	fmt.Println(len(bigPileArray))
}
