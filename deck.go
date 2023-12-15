package main
// package main indicates this file also belongs to the project main
// give multiple file names in the command to run both the files


import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)
//creating new type of deck 
// custom type declarations
type deck []string // like deck type extending slice of type str
// receiver fns i.e (d deck)
func (d deck) print() { // is like a method of instance deck i.e any var of type deck has access tothis method
	// d denoting deck as a varible
	// but the words this or self are not conventional to use in go
	// usually this var here d by convention is simply frst 1 or 2 letters of the type
	for index, card := range d { 
		fmt.Println(index, card)
	}
}

// create a deck of cards
func newDeck() deck{ //i.e retun type is of deck
	// no need of receiver function here because this method is probably called to intialise a new deck, so wont be passsing a var here
	cards := deck{} // intilise cards of type deck
	// add all the 52 cards manually 
	// or create two slices : one for card suites and other for values, then loop and add values to the dekc
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// when u've a variable, but no need to use, use an _
	for _, card := range cardSuites{
		for _, cardValue := range cardValues{
			cards = append(cards, card + " of " + cardValue)
		}
	}
	return cards
}


func deal(d deck, handSize int) (deck, deck) { // multiplereturn types
	// deal: ex deal of 3 i.e hadnsize of 3 cards, using slicing to split the cards
	return d[:handSize], d[handSize:]
}

// figure out reciver or arg
func (d deck) toString() string {
	//convert it to a [] str: 
	// []string(d)
	// then to str:
	// join the []str using a sep to a single string
	return strings.Join([]string(d), ",")
}


// write the current deck of cards to a local file
// make use of ioutil pkg, to write file, need to write in bytes format
func (d deck) saveToFile(filename string) error { //ioutil module returns ann error if occurs, we return the same
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 default permission anyone can read and write teh file
}


func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //returs bytesice and error if amy or nill
	if err!= nil {
		//option1: log the error and call newDeck so that fn still returns  a deck type
		//option2: log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) //any code other than 0 indicates the process is failed
	}
	bs_str := string(bs) //convert to str eg: "Ace of hearts, One of hearts" from bytes
	bs_slice_str :=  strings.Split(bs_str, ",") //convert to []string
	return deck(bs_slice_str)
}


func (d deck) shuffle_dup(){
	//generate random number in len of cards and swap between the curr card and random card
	// use rand module in math pkg, Intn func returns an int in the given int range
	for i := range d {
		newPosition := rand.Intn(len(d)-1)
		// but here's a case when all the last ele are same after every random generation, bcz the way how random geenrator works is generates a random value out of a seed i.e source
		// so the seed is always same and so the random number is also same 
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) shuffle(){
	// source := rand.NewSource() // new seed every time
	// send new source every time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // returns rand that uses random val from source to generate random values
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}