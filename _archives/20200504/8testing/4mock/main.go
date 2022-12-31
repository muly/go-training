package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Contact struct {
	Age int `csv:"age,omitempty"`
}

func process()int {
	c := []Contact{}
	err := loadCsv("data.csv", &c)
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for range c{
		sum+=c.Age
	}
	return sum
}


func process2(c []Contact{})int {
	// c := []Contact{}
	// err := loadCsv("data.csv", &c)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	sum := 0
	for range c{
		sum+=c.Age
	}
	return sum
}


func loadCsv(csvFile string, out interface{}) error {
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return gocsv.UnmarshalFile(f, out)
}
