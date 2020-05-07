package data

import (
	"fmt"

	"github.com/muly/go-training/20200504/9microservice/models"
)

// var customerIndexLoc = map[string]models.Customer{
// 	"india": models.Customer{Name: "C1", Address: "India"},
// 	"usa":   models.Customer{Name: "C2", Address: "USA"},
// 	"uk":    models.Customer{Name: "C3", Address: "UK"},
// }

func GetCustomersByFilter(loc string) models.Customer {

	dbConn, err := Open()
	if err != nil {
		fmt.Println(err)
		return models.Customer{}
	}
	

	rows, err := dbConn.Query(fmt.Sprintf("SELECT name, address from public.customer where address = '%s'", loc))
	if err != nil {
		fmt.Println(err)
		return models.Customer{}
	}

	c := models.Customer{}
	for rows.Next() {
		var name, address string
		err := rows.Scan( &name, &address)
		if err != nil {
			fmt.Println(err)
			return models.Customer{}
		}
		c.Name = name
		c.Address = address
	}
	return c
}
