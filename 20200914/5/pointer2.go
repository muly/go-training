// what is the output

// pointer 

package main
import("fmt")

type student struct {
	name          string
	email         string
	english       int
	maths         int
	total int
}

func cleanup(s *student){
	s = &student{}
	// OR
	// s = new(student)
	// OR
	// *s =  student{}
}

func main() {
	s := student{name: "student1", email: "dsdss"}
	s.english = 100
	s.maths = 90

	cleanup(&s)

	fmt.Println(s)
}
