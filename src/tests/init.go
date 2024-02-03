package tests

import (
	"Redis/redis-impl-go/src/app/repository/impl"
	implservice "Redis/redis-impl-go/src/app/service/impl"
	"Redis/redis-impl-go/src/app/utils/database"
)

var (
	db            = database.DBConnect()
	CourseRepo    = impl.NewCourseRepository(db)
	CourseService = implservice.NewCourseService(CourseRepo)
)
