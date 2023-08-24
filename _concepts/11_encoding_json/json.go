// TODO: need to add the structs concepts with examples

package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"addr"`
	Marks   int    `json:"marks,omitempty"`
}

func main() {
	var marks int
	fmt.Println("enter marks:")
	fmt.Scanf("%d", &marks)

	var s student = student{
		Name:    "rohit",
	}

	s.Marks = marks

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
