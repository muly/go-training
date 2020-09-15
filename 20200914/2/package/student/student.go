package student


type Student struct {
	Name          string
	Email         string
	English       int
	Maths         int
	Homeaddress   Address
	Schooladdress Address
	myPrivate string
}

// Address is the address struct to 
type Address struct {
	Line1, Line2 string
	City         string
	Zipcode      int
}