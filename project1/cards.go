package main 
import "fmt"
type deck []string
func (d deck) print(){//This is a method that takes a deck as a receiver and prints each card in the deck along with its index. The method uses a for loop to iterate over the deck, where i is the index and card is the value at that index. It then prints the index and the card using fmt.Println.
	for i, card := range d {
		fmt.Println(i, card)
	}
}