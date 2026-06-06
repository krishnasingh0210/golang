package main 
import "fmt"
type deck []string

func newDeck() deck {
cards:=deck{}
cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
for _, suit := range cardSuits { //using '_' means that we are ignoring the index of the loop and only using the value of the loop.
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
//RECEIVER FUNCTION
func (d deck) print(){//This is a method that takes a deck as a receiver and prints each card in the deck along with its index.
	for i, card := range d {
		fmt.Println(i, card)
	}
}
//DEAL FUNCTION
func deal(d deck, handsize int ) (deck,deck){
	return d[:handsize], d[handsize:]
}