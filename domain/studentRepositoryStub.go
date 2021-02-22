package domain

import "github.com/vipul-08/student-api/exceptions"

type StudentRepositoryStub struct {
	students []Student
}

func (repository StudentRepositoryStub) FindAll() ([]Student, *exceptions.AppError) {
	return repository.students, nil
}

func (repository StudentRepositoryStub) FindById(id int) (*Student, *exceptions.AppError) {
	for _, student := range repository.students {
		if student.Id == id {
			return &student, nil
		}
	}
	return &Student{}, nil
}

func NewStudentRepositoryStub() StudentRepositoryStub {
	students := []Student{
		{
			1001,
			"Vipul Singh Raghuvanshi",
			23,
			"BE",
			"CE",
		},
		{
			1002,
			"Vikram Parmar",
			22,
			"BE",
			"EXTC",
		},
	}
	return StudentRepositoryStub{students}
}
