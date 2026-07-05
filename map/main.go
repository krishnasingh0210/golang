package main
import ("fmt")

func main(){
	/*
	colors := make(map[string]string)//Both are the same.
	colors["white"] = "#ffffff"
	//var colors map[string]string
	//delete(colors, "white")//delete a key-value pair from the map
	*/
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	printMap(colors)
}
func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}