// user defined errors
package main

import "fmt"
import "errors"
import "log"

var divideByZero = errors.New("division by 0 not allowed")
var invalidData = errors.New("invalid data, p must be > q")

func main(){

	// 
	res, err := division(1,5)
	if err != nil{
		if errors.Is(err, divideByZero){
			fmt.Println("received a division by 0 error")
		}
		if errors.Is(err, invalidData){
			fmt.Println("received invalid data error")
		}
		log.Println(err.Error())
		return 
	}

	fmt.Println(res)

}

func division(p, q float32)(float32, error){
	if q==0{
		return 0, divideByZero
	}


	if p < q {
		return 0, invalidData
	}

	return p/q, nil
}