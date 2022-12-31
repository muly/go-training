package pk2

import (
	"fmt"

	"github.com/muly/go-training/2022/1/7packages3-cyclic-dependency/pk1"
)

var J = 10

func something() {
	fmt.Println(pk1.I)
}
