package main

import "fmt"

func main() {
	var color string = "RED"

	if color == "RED" || color == "BLUE" {
		fmt.Println("my fav color")
	} else {
		fmt.Println("not my fav color")
	}

	switch color {
	case "BLUE":
		fmt.Println("my fav color")
	case "RED":
		fmt.Println("my fav color")
	default:
		fmt.Println("not my fav color")
	}

	switch color {
	case "BLUE", "RED":
		fmt.Println("my fav color")
	default:
		fmt.Println("not my fav color")
	}
}
