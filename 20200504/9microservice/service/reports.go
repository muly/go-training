package service

import (
	"github.com/muly/go-training/20200504/9microservice/client"
	"github.com/muly/go-training/20200504/9microservice/data"
	"github.com/muly/go-training/20200504/9microservice/models"
)

func GetReports(loc string) (models.Report, error) {
	// add business logic here

	cust, err := client.GetCustomer(loc)
	if err != nil {
		return models.Report{}, err
	}

	c := data.GetReports(cust)

	return c, nil
}
