package main // There can be only one main package in a Go program, and it must contain the main function, which is the entry point of the program.

import "fmt"

func add() {
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Println("The numbers are:", a+b)
}
//STRING REVERSE FUNCTION

func string_reverse() {
	var txt, rev string
	fmt.Println("Enter a string:")
	fmt.Scan(&txt)
	for i := len(txt) - 1; i >= 0; i-- {
		rev += string(txt[i])
	}
	fmt.Println("Reversed string:", rev)
}
//CARDS FUNCTION

func cards() {
  deckSize := 52 //initializing a variable (if want to write this outside write var deckSize int = 52)
  fmt.Println(deckSize)
}
//CARD SLICE FUNCTION

func cardslice(){
	  deck := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	  deck= append(deck, "Four of Clubs")
	  fmt.Println(deck)
}
//NEW CARD FUNCTION

func newcard() string {
  return "Ace of Spades"
}
//COLOR FUNCTION

func colors() (string, string, string) {
return "red", "yellow", "blue"
}
//MAIN FUNCTION
func main(){
    add()
    string_reverse()
    cards()
    cardslice()
	fmt.Println(newcard())

    cardds := newDeck()
	cardds.print()
	/*cardds := deck{"Ace of Hearts", "Two of Diamonds", "Three of Clubs"}//creating a new type of deck and initializing it with some values
	cardds.print()*/

	hand, remainingdeck := deal(cardds,5)/*this means that cards from cardds till (5-1) and the value of type deck will be assigned to hand 
	and the remaining cards from cardds will be assigned to remainingdeck.*/
	hand.print()
	remainingdeck.print()

    color1, color2, color3 := colors()
    fmt.Println(color1, color2, color3)
}

