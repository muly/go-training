package pk1

import (
	"fmt"

	"github.com/muly/go-training/2022/1/7packages3-cyclic-dependency/pk2"
)

var I = 10

func something() {
	fmt.Println(pk2.J)
}
