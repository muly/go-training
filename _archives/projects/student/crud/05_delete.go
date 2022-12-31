package crud

import (
	"fmt"
)

// Delete deletes the record
func Delete(id string) error {
	if id == "" {
		return fmt.Errorf("key fields are missing: %w", ErrorInvalidData)
	}
	exists, _ := existsById(id)
	if !exists {
		return fmt.Errorf("%w : %v", ErrorNotFound, id)
	}

	i := -1
	for i = range AllStudents {
		if AllStudents[i].Id == id {
			break
		}
	}

	AllStudents = append(AllStudents[:i], AllStudents[i+1:]...)

	return nil
}
