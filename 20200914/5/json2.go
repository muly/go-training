// with nested structs
// omitempty
// -

package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Color string  `json:"color"`
	Breed string `json:"breed,omitempty"`
	Owner owner1  `json:"owner"`
	MyAge   int    `json:"-"` 
}

type owner1 struct {
	OwnerName string `json:"owner_name,omitempty"`
	City      string `json:"city,omitempty"`
}

func main() {
	
	unmarshalExample()
}


func unmarshalExample() {
	jsonStr := `{"color":"brown","breed":"","age":5,"owner":{"owner_name":"ABC","city":"Blore"}}`
	d := dog{}
	 json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(d.Owner) // 


	b, _ := json.Marshal(&d)
	fmt.Println(string(b))
}

