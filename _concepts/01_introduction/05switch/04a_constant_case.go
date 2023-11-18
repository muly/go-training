package main

import "fmt"

func main() {
	var color string = "RED"

	switch {
	case color == "BLUE":
		fmt.Println("violets are blue")
	case color == "RED":
		fmt.Println("roses are red")
	case color == "YELLOW":
		fmt.Println("sunflower is yellow")
	case color == "WHITE":
		fmt.Println("snow is white")
	case color == "BLACK":
		fmt.Println("night sky is black")
	default:
		fmt.Println("sorry this color is not supported by our program")
	}
}
