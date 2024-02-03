package response

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/model/web"
)

func ToCourseResponse(course domain.Course) web.CourseResponse {
	return web.CourseResponse{
		Id:          course.Id,
		AuthorId:    course.AuthorId,
		Title:       course.Title,
		Slug:        course.Slug,
		Description: course.Description,
		Perks:       course.Perks,
		Price:       course.Price,
		Banner:      course.Banner,
	}
}

func ToCourseResponses(courses []domain.Course) []web.CourseResponse {
	courseResponses := []web.CourseResponse{}
	for _, course := range courses {
		courseResponses = append(courseResponses, ToCourseResponse(course))
	}

	return courseResponses
}
