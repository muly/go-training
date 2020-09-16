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


func (a address) fullAddress()string{
	return fmt.Sprintf("%s\n%s\n%s\n%d\n", a.line1, a.line2, a.city, a.zipcode)
}


func main() {
	s := student{name: "student1", email: "dsdss"}

	s.english = 100
	s.city = "Bangalore" // accessing the embedded struct fields directly on parent struct
	fmt.Println(s.address.fullAddress())
	fmt.Println(s.fullAddress()) // accessing the embedded struct methods directly on parent struct

	fmt.Println(s)
}
