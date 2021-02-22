package domain

import (
	"github.com/vipul-08/student-api/exceptions"
)

// gorm.Model
type Student struct {
	Id     int    `json:"id" xml:"id" gorm:"primary_key;auto_increment"`
	Name   string `json:"name" xml:"name"`
	Age    uint8  `json:"age" xml:"age"`
	Class  string `json:"class" xml:"class"`
	Branch string `json:"branch" xml:"branch"`
}

// swagger:parameters getStudent deleteStudent
type studentIdParameter struct {
	// The id of a specific student
	// in: path
	// required: true
	ID int `json:"id"`
}

// A list of students returns in response
// swagger:response studentsResponse
type studentsResponse struct {
	// All students in the DB
	// in: body
	Body []Student
}

// swagger:response noContent
type studentNoContent struct {}


// A student object returns in response
// swagger:response studentResponse
type studentResponse struct {
	// Specific student in the DB
	// in: body
	Body Student
}

// swagger:parameters updateStudent createStudent
type studentParams struct {
	// Student data structure to Update or Create.
	// Note: the id field is ignored by update
	// in: body
	// required: true
	Body Student
}

//go:generate mockgen -destination=../mocks/domain/mock_StudentRepository.go -package=domain github.com/vipul-08/student-api/domain StudentRepository
type StudentRepository interface {
	FindAll() ([]Student, *exceptions.AppError)
	FindById(id int) (*Student, *exceptions.AppError)
	Insert(student *Student) (*Student, *exceptions.AppError)
	Update(student *Student) (*Student, *exceptions.AppError)
	Delete(id int) (int64, *exceptions.AppError)
}
