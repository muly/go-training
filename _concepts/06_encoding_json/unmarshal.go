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
	data := `{"name":"rohit","addr":"india","marks":300}`

	var s student

	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", s)

}
