package service

import (
	"Redis/redis-impl-go/src/app/model/web"
	"Redis/redis-impl-go/src/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	input := web.CourseCreateInput{
		AuthorId:    28,
		Title:       "Belajar Go",
		Slug:        "belajar-go",
		Description: "Belajar Go Description",
		Perks:       "Perks-1, Perks-2, Perks-3",
		Price:       99000,
		Banner:      "go.png",
	}

	courseResponse := tests.CourseService.Create(input)
	assert.NotNil(t, courseResponse.Id)
}

func TestGetAll(t *testing.T) {
	courseResponses := tests.CourseService.FindAll()
	assert.NotNil(t, courseResponses)
	assert.Equal(t, 2, len(courseResponses))
}
