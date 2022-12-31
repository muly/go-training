package main
import "fmt"


type math interface {
	isEven(y int) bool
	isOdd(y int) bool
}



type basicmath struct {
	name string
}

func (b basicmath) isEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func (b basicmath) isOdd(x int) bool {
	if x%2 != 0 {
		return true
	} else {
		return false
	}
}

func (b basicmath) double(x int) int {
	return x*2
}

func main() {
	var m math
	m = basicmath{"my basic math"}
	nbr :=10
	fmt.Println(nbr , m.isEven(nbr))
}
