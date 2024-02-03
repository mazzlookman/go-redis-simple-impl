package repository

import "Redis/redis-impl-go/src/app/model/domain"

type CourseRepository interface {
	Save(course domain.Course) domain.Course
	FindAll() []domain.Course
}
