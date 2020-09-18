// error wrapping
package main

import "fmt"
import "errors"
import "log"

var divideByZero = errors.New("division by 0 not allowed")
var invalidData = errors.New("invalid data, p must be > q")

func main() {
	//
	err := otherFunction()
	if err != nil {
		if errors.Is(err, divideByZero) {
			fmt.Println("received a division by 0 error")
		} else if errors.Is(err, invalidData) {
			fmt.Println("received invalid data error")
		} else {
			fmt.Println("unknown error")
		}

		log.Println(err.Error())
		return
	}

}

func otherFunction() error {
	///

	if _, err := division(1, 0); err != nil {
		// return err
		return fmt.Errorf("received error from division function: %w", err)
	}

	//

	return nil
}

func division(p, q float32) (float32, error) {
	if q == 0 {
		return 0, divideByZero
	}

	if p < q {
		return 0, invalidData
	}

	return p / q, nil
}
