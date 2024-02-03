package tests

import (
	implrepository "Redis/redis-impl-go/src/app/repository/impl"
	implservice "Redis/redis-impl-go/src/app/service/impl"
	"Redis/redis-impl-go/src/app/utils/database"
)

var (
	db            = database.DBConnect()
	CourseRepo    = implrepository.NewCourseRepository(db)
	CourseService = implservice.NewCourseService(CourseRepo)
)
