package impl

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/model/web"
	"Redis/redis-impl-go/src/app/repository"
	"Redis/redis-impl-go/src/app/service"
	"Redis/redis-impl-go/src/app/utils/errors"
	"Redis/redis-impl-go/src/app/utils/response"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type CourseServiceImpl struct {
	repository.CourseRepository
	redisClient *redis.Client
}

var ctx = context.Background()

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

	courseResponses := response.ToCourseResponses(courses)

	// save to redis
	marshal, _ := json.Marshal(&courseResponses)

	_, err := s.redisClient.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.Set(ctx, "courses", string(marshal), 0)
		pipeliner.Set(ctx, "is_save", true, 0)

		return nil
	})

	errors.PanicIfError(err)

	return courseResponses
}

func (s *CourseServiceImpl) IsSaveToRedis() bool {
	isSave := s.redisClient.Get(ctx, "is_save").Val()
	if isSave == "1" {
		return true
	}
	return false
}

func (s *CourseServiceImpl) FindAllWithRedis() []web.CourseResponse {
	courseResponses := []web.CourseResponse{}

	courses := s.redisClient.Get(ctx, "courses").Val()

	byteCourses := []byte(courses)
	err := json.Unmarshal(byteCourses, &courseResponses)
	errors.PanicIfError(err)

	return courseResponses
}

func NewCourseService(courseRepository repository.CourseRepository, rcli *redis.Client) service.CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository, redisClient: rcli}
}
