package main

// SEE "import cycle error.svg"

// // scenario 1 // no error
// pk1 imports pk2
// pk3 imports pk2


// // scenario 2 // ERROR: import cycle not allowed
// pk1 imports pk2
// pk2 imports pk1

// pk1 -> pk2 -> pk1

// // scenario 3 // ERROR: import cycle not allowed
// pk1 imports pk2
// pk2 imports pk3
// pk3 imports pk1

// pk1 -> pk2 -> pk3 -> pk1

func main(){

}