package main

import "fmt"
import "errors"
import "log"

func main(){

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := f1()
	if err != nil{
		log.Println(err)
		return 
	}


}


func f1()error{
	err := f2()
	if err != nil{
		return err
	}
	return nil
}

func f2()error{
	err := f3()
	if err != nil{

		return err
	}
	return nil
}

func f3()error{
	res, err := division(1,0)
	if err != nil{
		return fmt.Errorf("division() error: %w", err)
	}

	fmt.Println(res)
	return nil
}



func division(p, q float32)(float32, error){
	if q==0{
		return 0, errors.New("division by 0 not allowed")
	}
	return p/q, nil
}