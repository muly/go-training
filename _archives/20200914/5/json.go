package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Colour string  // should match (except the starting captial letter), else json tag should be specified 
	breed string `json:"breed"` // should be exported field
	Owner string // will work
	MyAge   int    `json:"age"` // will work
}

func main() {
	// marshalExample()
	unmarshalExample()
}

// func marshalExample() {
// 	d := dog{Colour: "brown", Breed: "German Shepherd", Age: 5}
// 	b, _ := json.Marshal(&d)
// 	fmt.Println(string(b))
// }

func unmarshalExample() {
	jsonStr := `{"color":"brown","breed":"German Shepherd","age":5, "owner": "abc"}`
	d := dog{}
	 json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(d) // {"", "German Shepherd", 5}
}




	// read
	// func json.Unmarshal(data []byte, v interface{}) error
	// json.NewDecoder(r io.Reader).Decode(v interface{})


	// write
	// func json.Marshal(v interface{}) ([]byte, error)
	// json.NewEncoder(w io.Writer).Encode(v interface{}) error

	
//// Marshal: struct to json string
//// Unmarshal: json string to struct


//// encode: from struct
//////	to stdout
//////	to json file
//////	to http ResponseWriter as json
//// decode: to struct
//////	from a json string
//////	from a json file
//////	from json data in http request


// examples: 
// 		https://github.com/muly/howto/blob/master/golang/encoding/json/encode-decode-example.go
// 		https://github.com/muly/howto/blob/master/golang/encoding/json/marshal-unmarshal-example.go





// Questions: 
// 	s student{"A"} -> s.json  // encode
// 	`{"name":"A"}` -> student{} // unmars