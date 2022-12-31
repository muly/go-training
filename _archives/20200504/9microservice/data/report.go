package data

import (
	"github.com/muly/go-training/20200504/9microservice/models"
)

var allRecordsIndexCust = map[string]models.Report{
	"C1": models.Report{Customer: "C1", ReportName: "R1", AverageCost: 100},
	"C2": models.Report{Customer: "C2", ReportName: "R2", AverageCost: 200},
	"C3": models.Report{Customer: "C3", ReportName: "R3", AverageCost: 300},
}

func GetReports(cust models.Customer) models.Report {
	return allRecordsIndexCust[cust.Name]
}
