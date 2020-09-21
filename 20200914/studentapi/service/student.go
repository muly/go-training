package service

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/muly/go-training/20200914/studentapi/model"
)

func GetStudent(session *mgo.Session, id string) (model.Student, error) {
	ms := model.Student{}

	filter := bson.D{{Name: "_id", Value: id}}
	err := session.DB("local").C("student").Find(filter).One(&ms)
	if err != nil {
		return model.Student{}, err
	}

	return ms, nil
}
