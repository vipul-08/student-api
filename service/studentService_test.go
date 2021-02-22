package service

import (
	"github.com/golang/mock/gomock"
	mainDomain "github.com/vipul-08/student-api/domain"
	"github.com/vipul-08/student-api/exceptions"
	"github.com/vipul-08/student-api/mocks/domain"
	"testing"
)

var mockRepository *domain.MockStudentRepository
var service StudentService

func setup(t *testing.T) func()  {
	mockController := gomock.NewController(t)
	mockRepository = domain.NewMockStudentRepository(mockController)
	service = NewStudentService(mockRepository)
	return func() {
		service = nil
		defer mockController.Finish()
	}
}

func Test_create_student_successfully(t *testing.T) {
	// Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	payloadStudent := mainDomain.Student{
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}

	studentWithId := payloadStudent
	studentWithId.Id = 999

	mockRepository.EXPECT().Insert(&payloadStudent).Return(&studentWithId, nil)
	newStudent,appError := service.InsertStudent(&payloadStudent)

	//Assert
	if appError != nil {
		t.Error("Test failed while creating account")
	}
	if newStudent.Id != studentWithId.Id {
		t.Error("Failed while matching new student id")
	}
}

func Test_create_student_failed(t *testing.T) {
	// Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	payloadStudent := mainDomain.Student{
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}

	mockRepository.EXPECT().Insert(&payloadStudent).Return(nil, exceptions.NewNotFoundError("Unable to Add"))
	_,appError := service.InsertStudent(&payloadStudent)

	//Assert
	if appError == nil {
		t.Error("Test failed while creating student")
	}
}

func Test_get_student_failed(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	mockRepository.EXPECT().FindById(10).Return(nil, exceptions.NewNotFoundError("Student Not Found"))
	_,appError := service.GetStudentById(10)

	//Assert
	if appError == nil {
		t.Error("Test failed while getting specific student")
	}
}

func Test_get_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	student := mainDomain.Student{
		Id: 	11,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().FindById(11).Return(&student, nil)
	_,appError := service.GetStudentById(11)

	//Assert
	if appError != nil {
		t.Error("Test failed while getting specific student")
	}
}

func Test_update_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	student := mainDomain.Student{
		Id: 	11,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().Update(&student).Return(&student, nil)
	_,appError := service.UpdateStudent(&student)

	//Assert
	if appError != nil {
		t.Error("Test failed while updating specific student")
	}
}

func Test_update_student_failure(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	student := mainDomain.Student{
		Id: 	11,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().Update(&student).Return(nil, exceptions.NewNotFoundError("Unable to Update"))
	_,appError := service.UpdateStudent(&student)

	//Assert
	if appError == nil {
		t.Error("Test failed while updating specific student")
	}
}

func Test_delete_student_failure(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	mockRepository.EXPECT().Delete(11).Return(int64(0), exceptions.NewNotFoundError("No rows affected"))
	_,appError := service.DeleteStudent(11)

	//Assert
	if appError == nil {
		t.Error("Test failed while deleting specific student")
	}
}

func Test_delete_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	mockRepository.EXPECT().Delete(11).Return(int64(1), nil)
	_,appError := service.DeleteStudent(11)

	//Assert
	if appError != nil {
		t.Error("Test failed while deleting specific student")
	}
}
