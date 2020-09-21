// Mysql Go example: not working
package main

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dbConn *gorm.DB

var (
	user     string = "root"
	secret   string = "password"
	dbip     string = "localhost"
	dbport   string = "3306"
	dbschema string = "dev"
)

func init() {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(user, ":", secret, "@tcp(", dbip, ":", dbport, ")/", dbschema))
	log.Println("MySQL Database Connection String :", buffer.String())
	dbURL := buffer.String()

	var err error
	dbConn, err = gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	dbConn.DB()
	err = dbConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(20)
	dbConn.SingularTable(true)
	dbConn.LogMode(true)
	return

}


type student struct {
	Name          string `gorm:"column:name"  json:"name"`
	Email         string `gorm:"column:email"  json:"email"`
	English       int `gorm:"column:english"  json:"english"`
	Maths         int `gorm:"column:maths"  json:"maths"` 
	Total int `gorm:"column:total"  json:"total"`
}


func main() {

	fmt.Println("Hello, playground")

	cust := Customer{Name: "test 1", Email: "test email 1", English: 70, Maths: 90}
	dbConn.Save(&cust)

	cust := Customer{Name: "test 2", Email: "test email 2", English: 75, Maths: 95}
	dbConn.Save(&cust)


	c := []Customer{}
	dbConn.Table("student").Scan(&c)
	fmt.Println("Active Customers list: ", c)
}

/*
create database dev;

use dev;

drop table student;

create table student(
name varchar(100),
email varchar(100),
english int,
maths int,
total int
);


*/