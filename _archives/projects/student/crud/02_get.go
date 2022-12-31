package crud

import (
	"errors"
	"fmt"

	data "github.com/muly/go-training/projects/student/db"
)

// ErrorInvalidData errors
var (
	ErrorInvalidData = errors.New("Invalid Data")
	ErrorNotFound    = errors.New("Not Found")
)

// Get gets the records based on the keys and their values provided
// TODO: func Get(filters url.Values) (data.Students, error) {}
func Get() (data.Students, error) {
	return AllStudents, nil

}

// GetById gets the record based on the id provided
func GetById(id string) (data.Students, error) {
	results := data.Students{}
	if id == "" {
		return data.Students{}, fmt.Errorf("key fields are missing: , %w", ErrorInvalidData)
	}

	found, results := existsById(id)
	if !found {
		return data.Students{}, fmt.Errorf("%w", ErrorNotFound)
	}

	return results, nil
}

func existsById(id string) (bool, data.Students) {
	results := data.Students{}
	flag := true

	for i := range AllStudents {
		if AllStudents[i].Id == id {
			results = append(results, AllStudents[i])
		}
	}
	if len(results) == 0 {
		flag = false
	}

	return flag, results
}
