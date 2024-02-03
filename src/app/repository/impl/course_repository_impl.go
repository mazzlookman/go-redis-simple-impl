package impl

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/repository"
	"Redis/redis-impl-go/src/app/utils/errors"
	"gorm.io/gorm"
)

func NewCourseRepository(DB *gorm.DB) repository.CourseRepository {
	return &CourseRepositoryImpl{DB: DB}
}

type CourseRepositoryImpl struct {
	*gorm.DB
}

func (r *CourseRepositoryImpl) Save(course domain.Course) domain.Course {
	err := r.DB.Create(&course).Error
	errors.PanicIfError(err)

	return course
}

func (r *CourseRepositoryImpl) FindAll() []domain.Course {
	courses := []domain.Course{}
	err := r.DB.Find(&courses).Error
	errors.PanicIfError(err)

	return courses
}
