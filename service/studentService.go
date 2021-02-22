package service

import (
	"github.com/vipul-08/student-api/domain"
	"github.com/vipul-08/student-api/exceptions"
)

//go:generate mockgen -destination=../mocks/service/mock_StudentService.go -package=service github.com/vipul-08/student-api/service StudentService
type StudentService interface {
	GetAllStudents() ([]domain.Student, *exceptions.AppError)
	GetStudentById(id int) (*domain.Student, *exceptions.AppError)
	InsertStudent(s *domain.Student) (*domain.Student, *exceptions.AppError)
	UpdateStudent(s *domain.Student) (*domain.Student, *exceptions.AppError)
	DeleteStudent(id int) (int64, *exceptions.AppError)
}

type DefaultStudentService struct {
	repo domain.StudentRepository
}

func (service DefaultStudentService) GetAllStudents() ([]domain.Student, *exceptions.AppError) {
	return service.repo.FindAll()
}

func (service DefaultStudentService) GetStudentById(id int) (*domain.Student, *exceptions.AppError) {
	return service.repo.FindById(id)
}

func (service DefaultStudentService) InsertStudent(s *domain.Student) (*domain.Student, *exceptions.AppError) {
	return service.repo.Insert(s)
}

func (service DefaultStudentService) UpdateStudent(s *domain.Student) (*domain.Student, *exceptions.AppError) {
	return service.repo.Update(s)
}

func (service DefaultStudentService) DeleteStudent(id int) (int64, *exceptions.AppError) {
	return service.repo.Delete(id)
}

func NewStudentService(repository domain.StudentRepository) DefaultStudentService {
	return DefaultStudentService{repository}
}
