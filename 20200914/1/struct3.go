package main
import("fmt")

type student struct {
	name          string
	email         string
	english       int
	maths         int
	address
}
type address struct {
	line1, line2 string
	city         string
	zipcode      int
}

func main() {
	s := student{name: "student1", email: "dsdss"}

	s.english = 100
	s.city = "Bangalore"

	fmt.Println(s)
}
