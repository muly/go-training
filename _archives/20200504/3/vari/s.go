package variable

type Student struct {
	Fname string
	Lname string
	marks int
}

// public
func (s Student) FullName() string {
	return s.fullName()
}

// private
func (s Student) fullName() string {
	return s.Fname + " " + s.Lname
}
