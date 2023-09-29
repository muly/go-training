package main

// START OMIT
type shape interface {
	area() float32
}

func (s string) area() float32 { // ERROR: cannot define new methods on non-local type string
	return float32(len(s) * len(s))
}

// END OMIT

func main(){
	
}
