package main 
import (
	"fmt"//importing the fmt package to use the Println and Scan functions
	"strings"//importing the strings package to use the strings.Join function
	"io/ioutil"//importing the ioutil package to use the WriteFile function
	"os"//importing the os package to use the Exit function
	"math/rand"//importing the math/rand package to use the Intn function
)
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
//Converting a deck to a string
func (d deck) toString() string {
	//[]string(d) converting the deck to a slice of strings
	return strings.Join([]string(d), ", ")
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)//0666 is the default code perimission means that the file can be read and written by anyone.
}
func newDeckFromFile(filename string) deck {
	bs,err := ioutil.ReadFile(filename)//on the place of _ there is bs for byteslice but it is not used in this function so we are using _ to ignore it.
	if err !=nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")//converting the byteslice to a string
	return deck(s)
}
func (d deck) shuffle(){
	for i :=range d{
		newPosition := rand.Intn(len(d)-1)
		d[i],d[newPosition] = d[newPosition],d[i]
	}
}