package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	f, err := os.Open("testfile.txt")
	if err != nil{
		log.Println("could not open file:", err)
		return
	}
	defer f.Close()

	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}

	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}


	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}

	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}


	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}


	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}



	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}



	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}



	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}


	err = division(1, 0)
	if err != nil{
		log.Println("division error: ", err)
		return
	}



}




func division(p, q float32)(float32, error){
	if q==0{
		return 0, errors.New("division by 0 not allowed")
	}
	return p/q, nil
}