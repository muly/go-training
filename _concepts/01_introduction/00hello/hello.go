package main

import "fmt"

// step 1: package level initializations are executed first

func init() {
	// step 2: init() function if it is available, is executed second

}

func main() {
	// step 3: main() function will be executed third
	fmt.Println("Hello Class")
}
