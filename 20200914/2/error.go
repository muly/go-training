package main

import "fmt"
// import "errors"
import "log"

func main(){

	// 
	res, err := division(1,0)
	if err != nil{
		log.Println(err)
		return 
	}

	fmt.Println(res)

}

func division(p, q float32)(float32, error){
	if q==0{
		return 0, errors.New("division by 0 not allowed")
	}
	return p/q, nil
}