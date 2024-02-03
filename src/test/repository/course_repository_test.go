package repository

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/repository/impl"
	"Redis/redis-impl-go/src/app/utils/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	db         = database.DBConnect()
	courseRepo = impl.NewCourseRepository(db)
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

	save := courseRepo.Save(course)
	assert.NotNil(t, save.Id)
}

func TestFindAll(t *testing.T) {
	courses := courseRepo.FindAll()
	assert.Equal(t, 1, len(courses))
}
