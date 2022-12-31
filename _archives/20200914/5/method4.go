// pointer receiver

package main
import("fmt")

type student struct {
	name          string
	email         string
	english       int
	maths         int
	total int
}

func (s *student) updateTotal(){
	s.total = s.maths+s.english
}

func main() {
	s := student{name: "student1", email: "dsdss"}

	s.english = 100
	s.maths = 90

	s.updateTotal()

	fmt.Println(s)
}
