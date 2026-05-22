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
func main(){
    add()
    string_reverse()
}
