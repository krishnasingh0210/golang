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
//EVEN ODD FUNCTION
func even_odd(){
    var a int
fmt.Println("Enter the number:")
fmt.Scanln(&a)
if a%2==0 {
    fmt.Println("The entered number is even")
}else{//IT IS NECCESARY TO HAVE ELSE WRITE LIKE (  } else {  ) in the same line otherwise it will give error.
    fmt.Println("The entered number is odd")
} 
}
//PRIME NUMBER FUNCTION
func prime_number(){
	 var a,count int
fmt.Println("Enter the number:")
fmt.Scanln(&a)
for i :=1;i<=a;i++{
    if a%i==0{
        count++
    }
}
    if count==2{
        fmt.Println("The entered number is Prime.")
    }else{
        fmt.Println("The entered number is Not Prime.")
    }
}
//VOWELS FUNCTION
func vowels(){
	var a string
      var vowels_count int
      fmt.Println("Enter the string:")
      fmt.Scan(&a)
      for _,char :=range a{
          switch char{
              case 'a','e','i','o','u','A','E','I','O','U':
              vowels_count++
             fmt.Printf("Found vowels: %c\n", char) //fmt.Println("Found vowels:",char) //{This will result in rune(int32) value bcz char is rune type}
          }
      }
      fmt.Println("The vovels in the string are:",vowels_count)
}
//FACTORIAL FUNCTION
func factorial() {
    var a,fact int
    fact=1
  fmt.Println("Enter the number:")
  fmt.Scan(&a)
  for i:=1;i<=a;i++{
      fmt.Print(i)
      fact*=i
  }
  fmt.Println()
  fmt.Println("The factorial of the entered number is:",fact)
}

//MAIN FUNCTION
func main(){
   /*1. add()
    2. string_reverse()
    3. cards()
    4. cardslice()
	fmt.Println(newcard())
	11. even_odd()
	12. prime_number()
	13. vowels()
	14. factorial()
	
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

