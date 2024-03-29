package repository

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	course := domain.Course{
		AuthorId:    28,
		Title:       "Belajar Java",
		Slug:        "belajar-java",
		Description: "Kursus Belajar Java",
		Perks:       "Perks-1, Perks-2, Perks-3",
		Price:       99000,
		Banner:      "java.png",
	}

	save := tests.CourseRepo.Save(course)
	assert.NotNil(t, save.Id)
}

func TestFindAll(t *testing.T) {
	courses := tests.CourseRepo.FindAll()
	assert.Equal(t, 1, len(courses))
}
