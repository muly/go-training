package crud

import (
	"encoding/json"
	"fmt"

	data "github.com/muly/go-training/projects/student/db"
)

// Patch updates the record with only provided fields
func Patch(id string, updates map[string]interface{}) (data.Student, error) {
	result := data.Student{}

	if id == "" {
		return data.Student{}, fmt.Errorf("key fields are missing: %w", ErrorInvalidData)
	}

	exists, out := existsById(id)
	if !exists {
		return data.Student{}, fmt.Errorf("%v %w", id, ErrorNotFound)
	}

	output := out[0]
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(output)
	if err != nil {
		return data.Student{}, err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return data.Student{}, err
	}

	// iterate through inrecs
	for field := range inInterface {
		for k, v := range updates {
			if field == k {
				inInterface[field] = v
			}
		}
	}

	one, err := json.Marshal(inInterface)
	if err != nil {
		return data.Student{}, err
	}
	err = json.Unmarshal(one, &result)
	if err != nil {
		return data.Student{}, err
	}
	err = Put(result)
	if err != nil {
		return data.Student{}, err
	}

	results, err := GetById(id) //TODO: change return type from data.Students to data.Student
	if err != nil {
		return data.Student{}, err
	}

	return results[0], nil
}
