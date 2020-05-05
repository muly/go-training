package variable

var Marks = 1000

func Swap(a, b int) (int, int) {
	return b, a
}

func Add(a, b int) (int, int) {
	sum := a + b

	avg := (a + b) / 2

	return sum, avg
}

// variadic func
func AddAll(name string, marks ...int) int {
	var sum int
	for _, m := range marks {
		sum = sum + m
	}

	return sum
}



type MyAdd func(a, b int) (int, int)

var F MyAdd = func(a, b int) (int, int) {
	sum := a + b
	avg := (a + b) / 2
	return sum, avg
}
