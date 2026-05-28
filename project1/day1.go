package main // There can be only one main package in a Go program, and it must contain the main function, which is the entry point of the program.

import "fmt"

func add() {
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Println("The numbers are:", a+b)
}
func string_reverse() {
	var txt, rev string
	fmt.Println("Enter a string:")
	fmt.Scan(&txt)
	for i := len(txt) - 1; i >= 0; i-- {
		rev += string(txt[i])
	}
	fmt.Println("Reversed string:", rev)
}
func cards() {
  deckSize := 52 //initializing a variable (if want to write this outside write var deckSize int = 52)
  fmt.Println(deckSize)
}
func cardslice(){
	  deck := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	  deck= append(deck, "Four of Clubs")
	  fmt.Println(deck)
}
func newcard() string {
  return "Ace of Spades"
}
func main(){
    add()
    string_reverse()
    cards()
    cardslice()
	fmt.Println(newcard())
}
