package main

func main() {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Introduction
	// maps data-structure holds the key value pairs

	// syntax: map[KeyType]ValueType

	// reference: 
	// Go maps in action: https://go.dev/blog/maps

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## declaring a map
	// 01declaring/

	// 1declareOnly.go
	// { 
	// 	
	// 	var m map[int]string
	//
	// 	fmt.Println(m)
	// }

	// 2declareButUninitialized.go
	// { 
	// 	
	// 	var m map[int]string // not initialized
	//
	// 	if m == nil {
	// 		fmt.Println("map in nil")
	// 	}
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 2 initializing a map
	// 02initializing/

	// // 1 declare and initialize in separate lines
	// // 1declareAndInitialize_inSeparateLines_Example1.go
	// {
	// 	var m map[int]string
	// 	m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	fmt.Println(m)
	// }

	// // another example: map of phone number to person name
	// 1declareAndInitialize_inSeparateLines_Example2.go
	// {
	// 	var m map[string]string
	// 	m = map[string]string{
	// 		"123-456-7890": "John",
	// 		"123-456-9999": "Max",
	// 	}
	//
	// 	fmt.Println(m)
	// }


	// initializing: declare and initialize in same line
	// 2declareAndInitialize_inSameLine.go
	// {
	// 	var m map[int]string = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	fmt.Println(m)
	// }

	// initializing: declare and initialize in same line, inferring datatype
	// 3declareAndInitialize_inSameLineInferringDatatype.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	fmt.Println(m)
	// }

	
	// initializing: short hand form
	// 4declareAndInitialize_inSameLineShortHandForm.go
	// {
	// 	m := map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	fmt.Println(m)
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


	// initializing: individual values, after initializing
	// 6declareAndInitialize_individualValues_afterInitializing.go
	// {
	// 	var m map[int]string = map[int]string{9: "nine"}
	//
	// 	m[1] = "one"
	//
	// 	fmt.Println(m)
	// }


	// initializing: individual values, after empty initializing
	// 7declareAndInitialize_individualValues_afterEmptyInitializing.go
	// {
	// 	var m map[int]string = map[int]string{}
	//
	// 	m[1] = "one"
	//
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## assigning
	// 03assigning/

	// assigning one map to another map
	// 1fromAnotherMap.go
	// {
	// 	var m map[int]string
	//
	// 	n := map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	m = n
	//
	// 	fmt.Println(m)
	// TODO: print the address of both maps
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## reinitializing
	// 04reinitializing/

	// reinitializing the map
	// 1withAnotherMap.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	m = map[int]string{10: "ten", 20: "twenty", 30: "thirty"}
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## accessing Individual Element of the map
	// 5accessingIndividualElement/


	// reading individual element:
	// syntax: map_var[key]
	// 1reading.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	// 	fmt.Println(m[1])
	//
	// 	k := 0
	// 	fmt.Println(m[k])
	// }


	// writing individual element
	// syntax: map_var[key]=new_value
	// 2writing.go
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
	// 3writing_overwriting.go
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	m[3] = "A" // key has to be unique, if not overwrite
	// 	fmt.Println(m[3])
	//
	// 	m[1] = "A" // value need not be unique. here key 3, 1 both has the same value "A"
	//
	// }


	// accessing individual element: with existing key with empty value
	// 4reading_existingKeyWithEmptyValue.go
	// {
	// 	var m = map[int]float32{1: 111, 3: 333, 9: 0}
	//
	// 	fmt.Println(m[9])
	// 	fmt.Println(m)
	// }


	// accessing individual element: with not existing key
	// 5reading_withNotExistingKey.go
	// {
	// 	var m = map[int]float32{1: 111, 3: 333}
	//
	// 	fmt.Println(m[9]) // no error. returns zero value for the value type, i.e. 0 for float32
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 6 find if the key already exists in the map
	// as we noticed in the previous example, both of the cases (mentioned below) return same value, (ie 0). so how to differenciate between these two cases
	// - case 1: key exists but with zero value
	// - case 2: key doesn't exist in the map
	//
	// map returns a optional second value, of type bool, which indicates if the given key exists or not
	// syntax:
	//  value, exists := map_variable[key]
	// 		if exists is true, given key is in the map_variable
	//  	if exists is false, given key is not in the map_variable

	//
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
	// ## 7 len() system functions
	// len() system functions
	// {
	// 	var m = map[int]float32{1: 10.5, 3: 30.5, 5: 0.0}
	// 	fmt.Println(len(m))
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ## 8 for Loop
	// range returns key and value
	//
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k, v := range m {
	// 		fmt.Printf("key %v, value %v\n", k, v)
	// 	}
	// }

	// ## 8 for Loop
	// key and value are optional: value only
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for _, v := range m {
	// 		fmt.Printf("value %v\n", v)
	// 	}
	// }

	// ## 8 for Loop
	// key and value are optional: key only
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k, _ := range m {
	// 		fmt.Printf("key %v, value %v\n", k, m[k])
	// 	}
	// }

	// ## 8 for Loop
	// key and value are optional: key only, without underscore
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	//
	// 	for k := range m {
	// 		fmt.Printf("key %v, value %v\n", k, m[k])
	// 	}
	// }

	// ## 8 for Loop
	// order of retrieval is not guaranteed
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
	// // ## 9 make() system function
	// {
	// 	var m map[int]int = make(map[int]int) // functionally identical to initializing with empty map
	// // same as var m map[int]int = map[int]int{}
	// 	fmt.Printf("map %v, with length %d\n", m, len(m))
	// 	if m != nil {
	// 		fmt.Println("map in not nil")
	// 	}
	// }

	// // ## 9 make() system function
	// // Note: length parameter of make() has no effect for map type
	// {
	// 	var m = make(map[int]int, 10)
	// 	fmt.Printf("map %v, with length %d\n", m, len(m))
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// // ## 10 delete() system function
	// // delete one kv pair
	// // syntax: delete(map_variable, key)
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

	// // ## 10 delete() system function
	// // delete all
	// {
	// 	var m = map[int]string{1: "one", 3: "three", 0: "zero"}
	// 	for k := range m {
	// 		delete(m, k)
	// 	}
	// 	fmt.Println(m)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// // ## 11 map of complex types
	// // value may be of any type, including other maps.
	// // key may be of any type that are comparable (see https://go.dev/ref/spec#Comparison_operators).
	// // map of structs
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

	// // ## 11 map of complex types
	// // map of structs: cleaner example
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
	// map is pointer type
	// Note: this is Advanced topic. prerequisites: pointers, functions
	// If a map isn’t a reference variable, what is it?: https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
	//
	// 	{
	// 		var m = map[int]int{1: 100, 2: 200}
	// 		fmt.Println(m)
	// 		insert(m, 3, 300) // note: we are passing the map variable without * (like we do in case of the pointer type),
	// 		fmt.Println(m) // still, the change made in the insert() function scope is visible outside the insert() function, i.e in main() function.
	// 		// this behaviour usually happend if we pass the varaible as pointer. but since map is a pointer type, we have the same behavious without using *
	// 	}
	//
	// }
	//
	// func insert(m map[int]int, k, v int) {
	// 	m[k] = v
}
