package service

import (
	"Redis/redis-impl-go/src/app/model/web"
	"Redis/redis-impl-go/src/tests"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

var ctx = context.Background()

func TestCreate(t *testing.T) {
	for i := 0; i < 10000; i++ {
		input := web.CourseCreateInput{
			AuthorId:    28,
			Title:       "Belajar Course-" + strconv.Itoa(i),
			Slug:        "belajar-course-" + strconv.Itoa(i),
			Description: "Belajar Course-" + strconv.Itoa(i) + " Description",
			Perks:       "Perks-1, Perks-2, Perks-3",
			Price:       99000,
			Banner:      "course-" + strconv.Itoa(i) + ".png",
		}

		courseResponse := tests.CourseService.Create(input)
		assert.NotNil(t, courseResponse.Id)
	}
}

func TestGetAll(t *testing.T) {
	courseResponses := tests.CourseService.FindAll()
	assert.NotNil(t, courseResponses)
	assert.Equal(t, 10000, len(courseResponses))
}

func TestGetAllWithRedis(t *testing.T) {
	isSaveToRedis := tests.CourseService.IsSaveToRedis()
	if isSaveToRedis {
		log.Println("Get data from redis")
		findAllWithRedis := tests.CourseService.FindAllWithRedis()

		log.Println(findAllWithRedis)
		return
	}

	log.Println(isSaveToRedis)
}
