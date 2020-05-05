//
package main

import (
	"fmt"
	// "errors"
)

func main() {

	q, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("cannot divide by 0")
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
