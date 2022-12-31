package crud

import (
	"errors"

	data "github.com/muly/go-training/projects/student/db"
)

// Put updates the record
func Put(student data.Student) error {
	updated := false

	for i := range AllStudents {
		if AllStudents[i].Id == student.Id {
			updated = true
			AllStudents[i] = student
			break
		}
	}

	if !updated {
		return errors.New("Not Found")
	}
	return nil
}
