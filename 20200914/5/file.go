package main

import "os"
import "fmt"
import "io"

func main() {

	f, err := os.Open("data")
	if err != nil {

		fmt.Println(err)
		return
	}

	for {
		data := make([]byte, 5)
		_, err = f.Read(data)
		if err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Printf(string(data))
	}

	fmt.Println()
}
