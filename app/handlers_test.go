package app

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/vipul-08/student-api/domain"
	"github.com/vipul-08/student-api/exceptions"
	"github.com/vipul-08/student-api/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var mockHandlers StudentHandlers
var mockService *service.MockStudentService

func setup(t *testing.T) func() {
	mockController := gomock.NewController(t)
	mockService = service.NewMockStudentService(mockController)
	mockHandlers = StudentHandlers{mockService}
	router = mux.NewRouter()

	return func() {
		router = nil
		defer mockController.Finish()
	}
}

func Test_get_all_students_test_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	dummyStudentsList := []domain.Student {
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
	mockService.EXPECT().GetAllStudents().Return(dummyStudentsList, nil)
	router.HandleFunc("/students", mockHandlers.getAllStudents).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_all_students_test_500(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	mockService.EXPECT().GetAllStudents().Return(nil, exceptions.NewUnexpectedError("Unexpected DB Error"))
	router.HandleFunc("/students", mockHandlers.getAllStudents).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	mockService.EXPECT().GetStudentById(10).Return(nil, exceptions.NewNotFoundError("Student Not Found"))
	router.HandleFunc("/students/{id}", mockHandlers.getStudentById).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_specific_student_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Id:     10,
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().GetStudentById(10).Return(&student, nil)
	router.HandleFunc("/students/{id}", mockHandlers.getStudentById).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_post_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().InsertStudent(&student).Return(nil, exceptions.NewNotFoundError("Unable to Add"))
	router.HandleFunc("/students", mockHandlers.insertStudent).Methods(http.MethodPost)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_post_specific_student_201(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	studentWithId := student
	studentWithId.Id = 100

	mockService.EXPECT().InsertStudent(&student).Return(&studentWithId, nil)
	router.HandleFunc("/students", mockHandlers.insertStudent).Methods(http.MethodPost)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func Test_put_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Id:		100,
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().UpdateStudent(&student).Return(nil, exceptions.NewNotFoundError("Unable to Update"))
	router.HandleFunc("/students", mockHandlers.updateStudent).Methods(http.MethodPut)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPut, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_put_specific_student_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Id:		100,
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().UpdateStudent(&student).Return(&student, nil)
	router.HandleFunc("/students", mockHandlers.updateStudent).Methods(http.MethodPut)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPut, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_delete_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	mockService.EXPECT().DeleteStudent(10).Return(int64(0), exceptions.NewNotFoundError("No rows affected"))
	router.HandleFunc("/students/{id}", mockHandlers.deleteStudent).Methods(http.MethodDelete)
	req,_ := http.NewRequest(http.MethodDelete, "/students/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_delete_specific_student_204(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	mockService.EXPECT().DeleteStudent(10).Return(int64(1), nil)
	router.HandleFunc("/students/{id}", mockHandlers.deleteStudent).Methods(http.MethodDelete)
	req,_ := http.NewRequest(http.MethodDelete, "/students/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNoContent {
		t.Error("Failed while testing the status code")
	}
}
