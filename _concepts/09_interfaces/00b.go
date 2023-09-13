// Note: interface implementation must be done by a user defined type
// because we cannot add methods on a system type like string

package main

type shape interface {
	area() float32
}

func (s string) area() float32 { // ERROR: cannot define new methods on non-local type string
	return s.length * s.length
}
