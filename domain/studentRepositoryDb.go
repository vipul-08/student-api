package domain

import (
	"github.com/vipul-08/student-api/db"
	"github.com/vipul-08/student-api/exceptions"
	"gorm.io/gorm"
	"log"
)

type StudentRepositoryDb struct {
	client *gorm.DB
}

func NewStudentRepositoryDb() StudentRepositoryDb {
	client := db.GetDb()
	return StudentRepositoryDb{client}
}

func (d StudentRepositoryDb) FindAll() ([]Student, *exceptions.AppError) {
	students := make([]Student, 0)
	err := d.client.Debug().Model(&Student{}).Find(&students).Error
	if err != nil {
		log.Println("Error while fetching from DB " + err.Error())
		return students, exceptions.NewUnexpectedError("Unexpected DB Error")
	}
	return students, nil
}

func (d StudentRepositoryDb) FindById(id int) (*Student, *exceptions.AppError) {
	var student Student
	err := d.client.Debug().Model(&Student{}).Where("id = ?",id).Take(&student).Error
	if gorm.ErrRecordNotFound == err {
		return nil, exceptions.NewNotFoundError("Student Not Found")
	} else if err != nil {
		return nil, exceptions.NewUnexpectedError("Unexpected DB Error")
	}
	return &student, nil
}

func (d StudentRepositoryDb) Insert(s *Student) (*Student, *exceptions.AppError) {
	err := d.client.Debug().Create(&s).Error
	if err != nil {
		return nil, exceptions.NewNotFoundError("Unable to Add")
	}
	return s,nil
}

func (d StudentRepositoryDb) Update(s *Student) (*Student, *exceptions.AppError) {
	err := d.client.Debug().Model(&Student{}).Where("id = ?", s.Id).Take(&Student{}).Omit("id").Updates(&s).Error
	if err != nil {
		return nil, exceptions.NewNotFoundError("Unable to Update")
	}
	return s,nil
}

func (d StudentRepositoryDb) Delete(id int) (int64, *exceptions.AppError) {
	result := d.client.Debug().Model(&Student{}).Where("id = ?", id).Take(&Student{}).Delete(&Student{})
	if result.Error != nil {
		return 0, exceptions.NewNotFoundError("No rows affected")
	}
	return result.RowsAffected, nil
}
