package impl

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/model/web"
	"Redis/redis-impl-go/src/app/repository"
	"Redis/redis-impl-go/src/app/service"
	"Redis/redis-impl-go/src/app/utils/response"
)

type CourseServiceImpl struct {
	repository.CourseRepository
}

func (s *CourseServiceImpl) FindAllWithRedis() []web.CourseResponse {
	return nil
}

func (s *CourseServiceImpl) Create(input web.CourseCreateInput) web.CourseResponse {
	course := domain.Course{}
	course.AuthorId = input.AuthorId
	course.Title = input.Title
	course.Slug = input.Slug
	course.Description = input.Description
	course.Perks = input.Perks
	course.Price = input.Price
	course.Banner = input.Banner

	save := s.CourseRepository.Save(course)

	return response.ToCourseResponse(save)
}

func (s *CourseServiceImpl) FindAll() []web.CourseResponse {
	courses := s.CourseRepository.FindAll()

	// save to redis

	return response.ToCourseResponses(courses)
}

func NewCourseService(courseRepository repository.CourseRepository) service.CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository}
}
