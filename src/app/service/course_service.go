package service

import "Redis/redis-impl-go/src/app/model/web"

type CourseService interface {
	Create(input web.CourseCreateInput) web.CourseResponse
	FindAll() []web.CourseResponse
	IsSaveToRedis() bool
	FindAllWithRedis() []web.CourseResponse
}
