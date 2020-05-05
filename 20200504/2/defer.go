package main

import (
	"fmt"
)

func main() {

	if err := open(); err != nil {
		return
	}
	defer close()

	if err := process1(); err != nil {
		// close()
		return
	}

	if err := process2(); err != nil {
		// close()
		return
	}

	// close()
}

func open() {}

func close() {}

func process1() error {}

func process2() error {}
