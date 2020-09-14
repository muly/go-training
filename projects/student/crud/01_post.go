package crud

import (
	"fmt"

	data "github.com/muly/go-training/projects/student/db"

	"github.com/google/uuid"
)

// AllStudents is a collection of students
var AllStudents data.Students

// Post posts the given list of records into the database collection
// TODO: returns list of errors for all the failed records
func Post(students data.Students) (data.Students, error) {
	temp := data.Students{}

	for _, student := range students {
		if student.Id != "" {
			exist, _ := existsById(student.Id)
			if exist {
				return nil, fmt.Errorf("record already exists with the given id %s", student.Id)
			}
		} else {
			id, err := uuid.NewUUID()
			if err != nil {
				return nil, fmt.Errorf("error while generating new uuid: %v", err)
			}
			student.Id = id.String()
		}
		temp = append(temp, student)
	}
	AllStudents = append(AllStudents, temp...)
	return temp, nil
}
