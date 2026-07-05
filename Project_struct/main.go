package main
import "fmt"

type contactInfo struct {
    email string 
    zipcode int
}
type Person struct {
    firstName string
    lastName string
    contact contactInfo //can also write only contactInfo
}

func main() {
    /*
// Three ways to create a struct
    var alex Person
    alex.firstName = "Alex"
    alex.lastName = "Anderson"
    // alex := Person{firstName: "Alex", lastName: "Anderson"}
        // alex := Person{"Alex", "Anderson"} // This is also valid, but not recommended
// if no value in alex then it will be default value of string which is empty string
    fmt.Printf("%+v" ,alex)
    */
    jim := Person{
        firstName: "Jim",
        lastName:  "Beam",
        contact: contactInfo{ // This is how you can create a nested struct
            email:   "jim.beam@example.com",
            zipcode: 12345, // comma is required in end of each line in struct literal
        },
    }
   // fmt.Printf("%+v", jim)
   jim.updateName("Jimmy")// This still prints jim because of pass by value.(use pointer and receiver to update the value)
  jim.print()
}
func (pointerToPerson *Person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}

//This is how to create a receiver function in struct
func (p Person) print() {
    fmt.Printf("%+v", p)
}

