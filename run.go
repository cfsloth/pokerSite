package main

import (
	"fmt"

	"./domain/entity"
	"./domain/value"
)

func main() {
	//fmt.Println("Hello World")
	//kingSpades := entity.Card{1, "King", "Spades"}
	//str, _ := json.Marshal(card)
	//kingDiamonds := entity.Card{2, "King", "Diamonds"}
	//hand := entity.Hand{kingSpades, kingDiamonds}
	//str, _ := json.Marshal(hand)
	//fmt.Println(string(str))

	//Testing making a deck
	var deck value.Deck
	deck = value.CreatePokerDeck()
	//str, _ = json.Marshal(deck.CardArray)
	//fmt.Println(string(str))

	//Test Shuufling a deck
	for i := 0; i < 10; i++ {
		deck = value.RiffleShuffleDeck(deck)
		fmt.Println(deck)
	}
	hands := []entity.Hand{entity.Hand{},entity.Hand{},entity.Hand{},entity.Hand{}}
	deck, hands = value.DealingHandsAndStack(deck,hands)
	fmt.Println(hands)
}
