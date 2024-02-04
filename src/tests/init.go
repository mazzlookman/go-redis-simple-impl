package tests

import (
	implrepository "Redis/redis-impl-go/src/app/repository/impl"
	implservice "Redis/redis-impl-go/src/app/service/impl"
	"Redis/redis-impl-go/src/app/utils/database"
)

var (
	Db            = database.MysqlConnect()
	Rdb           = database.RedisClient()
	CourseRepo    = implrepository.NewCourseRepository(Db)
	CourseService = implservice.NewCourseService(CourseRepo, Rdb)
)
