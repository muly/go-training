package main

func main() {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## declaring a map
	// 01declaring/

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## initializing a map
	// 02initializing/

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## assigning
	// 04assigning/



	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## accessing Individual Element of the map
	// 05accessingIndividualElement/




	// accessing individual element: with existing key with empty value
	// 2reading_existingKeyWithEmptyValue.go
	// {
	// 	var m = map[int]float32{1: 111, 3: 333, 9: 0}
	//
	// 	fmt.Println(m[9])
	// 	fmt.Println(m)
	// }

	// accessing individual element: with non existing key
	// 3reading_withNonExistingKey.go
	// {
	// 	var m = map[int]float32{1: 111, 3: 333}
	//
	// 	fmt.Println(m[9]) // no error. returns zero value for the value type, i.e. 0 for float32
	// 	fmt.Println(m)
	// }

	// writing individual element
	// syntax: map_var[key]=new_value
	// 4writing.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	m[0] = "A"
	// 	fmt.Println(m[0])
	//
	// 	m[0] = "B"
	// 	fmt.Println(m[0])
	//
	// }

	// overwriting individual element
	// key_value has to be unique
	// syntax: map_var[key]=new_value
	// 5writing_overwriting.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	m[3] = "A" // key has to be unique, if not overwrite
	// 	fmt.Println(m[3])
	//
	// 	m[1] = "A" // value need not be unique. here key 3, 1 both has the same value "A"
	// 	fmt.Println(m)
	// }

	// find if the key already exists in the map
	// as we noticed in the previous example, both of the cases (mentioned below) return same value, (ie 0). so how to differenciate between these two cases
	// - case 1: key exists but with zero value
	// - case 2: key doesn't exist in the map
	//
	// map returns a optional second value, of type bool, which indicates if the given key exists or not
	// syntax:
	//  value, exists := map_variable[key]
	// 		if exists is true, given key is in the map_variable
	//  	if exists is false, given key is not in the map_variable

	// 6keyExistsCheck.go
	// {
	// 	var m = map[int]float32{1: 10.5, 3: 30.5, 5: 0.0}
	//
	// 	v, ok := m[99]
	// 	fmt.Printf("value: %v, key exists: %v\n", v, ok)
	//
	// 	v, ok = m[5]
	// 	fmt.Printf("value: %v, key exists: %v\n", v, ok)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## un-initialized map / nil map

	// 2declareButUninitialized.go // TODO: should I keep this program here under "declaring" section or remove it as we are covering this under "undeclared map" section. mostly lets ignore this program.
	// {
	// 	var m map[int]string // not initialized
	// 	if m == nil {
	// 		fmt.Println("map in nil")
	// 	}
	// }

	// initializing: by individual values on a nil map
	// A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic;
	// 5declareAndInitialize_byIndividualValuesOnANilMap.go
	// {
	// 	var m map[int]string // nil map, i.e. not initialized
	//
	// 	m[1] = "one" // ERROR: panic: assignment to entry in nil map
	//
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 6 len() system functions
	// 06len/
	// len() system functions
	// {
	// 	var m = map[int]float32{1: 10.5, 3: 30.5, 5: 0.0}
	// 	fmt.Println(len(m))
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 7 for Loop
	// 07loop/
	// syntax: key_var, val_var:= range map_var{}
	//
	// range returns key and value;

	// 1withKeyValue.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k, v := range m {
	// 		fmt.Printf("key %v, value %v\n", k, v)
	// 	}
	// }

	// key and value are optional: value only
	// 2withValueOnly.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for _, v := range m {
	// 		fmt.Printf("value %v\n", v)
	// 	}
	// }

	// key and value are optional: key only
	// 3withKeyOnly.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k, _ := range m {
	// 		fmt.Printf("key %v, value %v\n", k, m[k])
	// 	}
	// }

	// key and value are optional: key only, without underscore
	// 4withKeyOnly_withoutUnderscore.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k := range m {
	// 		fmt.Printf("key %v, value %v\n", k, m[k])
	// 	}
	// }

	// order of retrieval is not guaranteed
	// 5orderNotGuaranteed.go
	// {
	// 	var m = map[int]int{}
	// 	for k := 0; k < 10; k++ {
	// 		m[k] = k
	// 	}
	//
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// // ## 8 make() system function
	// 08make/

	// 1make.go
	// {
	// 	var m map[int]int = make(map[int]int) // functionally identical to initializing with empty map
	// // same as var m map[int]int = map[int]int{}
	// 	fmt.Printf("map %v, with length %d\n", m, len(m))
	// 	if m != nil {
	// 		fmt.Println("map in not nil")
	// 	}
	// }

	// 2make_withLenParam.go
	// // Note: length parameter of make() has no effect for map type
	// {
	// 	var m = make(map[int]int, 10)
	// 	fmt.Printf("map %v, with length %d\n", m, len(m))
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// // ## 9 delete() system function
	// 09delete/
	// // delete one kv pair
	// // syntax: delete(map_variable, key)

	// 1singlePair.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	// 	fmt.Println(m)
	//
	// 	_, ok := m[3]
	// 	fmt.Println("before delete, key exists: ", ok)
	//
	// 	delete(m, 3)
	// 	_, ok = m[3]
	// 	fmt.Println("after delete, key exists: ", ok)
	// 	fmt.Println(m)
	// }

	// // delete all
	// 2allPairs.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	// 	for k := range m {
	// 		delete(m, k)
	// 	}
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// // ## 10 map of complex types
	// 10withComplexTypes/
	// // value may be of any type, including other maps.
	// // key may be of any type that are comparable (see https://go.dev/ref/spec#Comparison_operators).

	// // map of struct
	// 1mapOfStruct.go
	// {
	// 	type personUnique struct {
	// 		email       string
	// 		phoneNumber string
	// 	}
	// 	type personDetails struct {
	// 		name    string
	// 		address string
	// 	}
	// 	var m = map[personUnique]personDetails{} // initialize with empty map to avoid panic error
	// 	m[personUnique{
	// 		email:       "email1@gmail.com",
	// 		phoneNumber: "123-456-7890",
	// 	}] = personDetails{
	// 		name:    "john",
	// 		address: "usa",
	// 	}
	// 	m[personUnique{
	// 		email:       "email2@gmail.com",
	// 		phoneNumber: "123-456-2222",
	// 	}] = personDetails{
	// 		name:    "max",
	// 		address: "usa",
	// 	}
	//
	// 	fmt.Println(m)
	//
	// 	fmt.Println(m[personUnique{
	// 		email:       "email1@gmail.com",
	// 		phoneNumber: "123-456-7890",
	// 	}])
	// }

	// // map of struct: cleaner example]
	// 2mapOfStruct_cleanerExample.go
	// {
	// // types for keys and values
	// 	type personUnique struct {
	// 		email       string
	// 		phoneNumber string
	// 	}
	// 	type personDetails struct {
	// 		name    string
	// 		address string
	// 	}
	//
	// // variables for keys and values
	// 	p1 := personUnique{
	// 		email:       "email1@gmail.com",
	// 		phoneNumber: "123-456-7890",
	// 	}
	// 	p2 := personUnique{
	// 		email:       "email2@gmail.com",
	// 		phoneNumber: "123-456-2222",
	// 	}
	// 	pd1 := personDetails{
	// 		name:    "john",
	// 		address: "usa",
	// 	}
	//
	// 	var m = map[personUnique]personDetails{} // initialize with empty map to avoid panic error
	// 	m[p1] = pd1
	// 	m[p2] = pd1 // using the same value with different key
	//
	// 	fmt.Println(m)
	//
	// 	fmt.Println(m[p1])
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 11 map is pointer type
	// 11mapIsPointerType/
	// Note: this is Advanced topic. prerequisites: pointers, functions
	// If a map isn’t a reference variable, what is it?: https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
	//
	// 1.go	// file already created

}
