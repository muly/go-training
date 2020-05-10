package service

import (
	"github.com/muly/go-training/20200504/9microservice/data"
	"github.com/muly/go-training/20200504/9microservice/models"
)

func GetCustomers(location string) models.Customer {
	// add business logic here

	c := data.GetCustomersByFilter(location)

	return c
}
