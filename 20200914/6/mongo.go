package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

type student struct {
	ID string `bson:"_id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Email         string `bson:"email" json:"email"`
	English       int `bson:"english" json:"english"`
	Maths         int `bson:"maths" json:"maths"` 
	Total int `bson:"total" json:"total"`
}


func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("error dial: ", err)
	}
	studentcollection := session.DB("local").C("student")

	s:= student{ID: "4", Name: "test 4", Email: "test email 4", English: 70, Maths: 90}
	err = studentcollection.Insert(s)
	if err != nil {
		fmt.Println("error insert: ", err)
	}

	ms := []student{}

	filter :=bson.D{{Name:"_id", Value:"4"}}
	err = studentcollection.Find(filter).All(&ms)
	if err != nil {
		fmt.Println("error find all: ", err)
	}
	fmt.Println(ms)
}
