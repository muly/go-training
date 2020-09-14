package data

// Student struct holds marks obtained and his/her grade and pass status
type Student struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	English        int    `json:"english"`
	Maths          int    `json:"maths"`
	Social         int    `json:"social"`
	Science        int    `json:"science"`
	SecondLanguage int    `json:"secondlanguage"`
	Total          int    `json:"total"`
	Grade          int    `json:"grade"`
}

// Students is a slice of Students struct holds all the students in the class
type Students []Student
