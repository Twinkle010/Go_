package main

import "fmt"

func main() {
	// Variable intialisation and assigning
	// var deck string = "Deck of Cards"
	// cards := "Set of Cards" // : refers to intialisation
	// fmt.Println(deck)
	// if var declared and not used, throws an error
	// fmt.Println(cards)
	// array or slices in go
	// deck = []string{"Ace"} // i.e [] indicates data type is slices and {} indicates what is stored inside the slice
	// cards:=[]string{"Ace", newCard()} 
	// //append to slices
	// cards = append(cards, "Spade")
	// //iterate through slices
	// for index, card := range cards { // meaning for index, card in range of cards
	// 	// stores as index and value in slices
	// 	fmt.Println(index, card)
	// }
	//since deck type is created now 
	// cards:= deck{"ace", newCard()}
	// for index, card := range cards { // meaning for index, card in range of cards
	// 	// stores as index and value in slices
	// 	fmt.Println(index, card)
	// }
	// cards := newDeck()
	// // cards.print()
	// // no need to import fns since these are declared in same pkg
	// hand, _ := deal(cards, 5)
	// // cards.print()
	// hand.print()
	// print("*")
	// remaining_deck.print()
	// greeting := "Hey!!" //type conversion
	// fmt.Println([]byte(greeting)) //type we need to change(input data)
	// //deck to str, how?
	// //deck-> []str->str->[]byte
	// fmt.Println(hand.toString())
	// hand.saveToFile("handCards")
	cards := newDeckFromFile("handCards")
	cards.print()
	// cards_x := newDeckFromFile("handCards.thatdoesn'texist")
	// cards_x.print()
	fmt.Println("***")
	cards.shuffle()
	cards.print()
}

// funcitons
func newCard() string{ //defines a function names newCard that return data of type string
	return "Deck of Cards"
}