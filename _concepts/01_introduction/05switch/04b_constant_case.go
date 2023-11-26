package main

import "fmt"

func main() {
	var color  string = "RED"

	switch color{
	case "BLUE":
		fmt.Println("violets are blue")
	case "RED":
		fmt.Println("roses are red")
	case "YELLOW":
		fmt.Println("sunflower is yellow")
	case "WHITE":
		fmt.Println("snow is white")
	case "BLACK":
		fmt.Println("black night sky")
	default:
		fmt.Println("sorry this color is not supported by our program")
	}
}
