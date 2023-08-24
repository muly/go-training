package main

import "fmt"

type car interface {
	drive()
	playMusic()
}

type teslamodel3 struct {
	// specs of car
}
func (c teslamodel3) drive() {
	// tesla model 3 implementation of drive
}
func (c teslamodel3) playMusic() {
	// tesla model 3 implementation of play music
}

type hondaAccord struct {
	// specs of car
}
func (c hondaAccord) drive() {
	// honda Accord implementation of drive
}
func (c hondaAccord) playMusic() {
	// honda Accord implementation of play music
}

type incompleteCar struct{}
func (c incompleteCar) drive() {
	// honda Accord implementation of drive
}

func main() {

	var c car

	t := teslamodel3{}
	h := hondaAccord{}
	i := incompleteCar{}

	c = t

	c = h

	c = i // ERROR: cannot use i (variable of type incompleteCar) as car value in assignment: incompleteCar does not implement car (missing method playMusic)

	fmt.Println(c)

}
