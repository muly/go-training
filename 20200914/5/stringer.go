// stringer

// {student1 fname dsdss 100 0 Bangalore}
// name:student1 fname, email:dsdss, english:100, maths:0, city:Bangalore}

package main
import("fmt")

type student struct {
	name          string
	email         string
	english       int
	maths         int
	city string
}

func (s student)String()string{
	return fmt.Sprintf("name:%s, email:%s, english:%d, maths:%d, city:%s",s.name, s.email, s.english, s.maths, s.city)
}


func main() {
	s := student{name: "student1 fname", email: "dsdss"}

	s.english = 100
	s.city = "Bangalore"

	fmt.Println(s.String())

///
	str := fmt.Sprintf("my output is %v",s)
	fmt.Println(str)

	//
	fmt.Println(s)
}
