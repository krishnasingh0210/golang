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
   /*1. add()
    2. string_reverse()
    3. cards()
    4. cardslice()
	fmt.Println(newcard())
	
    5. color1, color2, color3 := colors()
    fmt.Println(color1, color2, color3)

	6. greeting := "Hello World"
	fmt.Println([]byte(greeting))//slice of bytes that represents the string "Hello World"*/
	
    /*10. cardds := newDeck()
	fmt.Println(cardds.toString())//converting the deck to a string and printing it*/
	
	
	//7. cardds.saveToFile("my_cards.txt")   //saving the deck to a file named "my_cards.txt"
	
	/*8. cardds := newDeckFromFile("my_cards.txt")//creating a new deck from the file "my_cards.txt" but this will not work because 
	the file is not present in the current directory so we need to create the file first and then run this code for now this will our given error.
	cardds.print()//printing the new deck created from the file */
	
	/*9.cardds := deck{"Ace of Hearts", "Two of Diamonds", "Three of Clubs"}//creating a new type of deck and initializing it with some values
	cardds.print()*/
	
	 cardds := newDeck()
	 cardds.shuffle()
	cardds.print()

	hand, remainingdeck := deal(cardds,5)/*this means that cards from cardds till (5-1) and the value of type deck will be assigned to hand 
	and the remaining cards from cardds will be assigned to remainingdeck.*/
	hand.print()
	remainingdeck.print()
	
}

