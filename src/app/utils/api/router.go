package api

import (
	implcontroller "Redis/redis-impl-go/src/app/controller/impl"
	"Redis/redis-impl-go/src/app/repository/impl"
	implservice "Redis/redis-impl-go/src/app/service/impl"
	"Redis/redis-impl-go/src/app/utils/database"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	router := fiber.New()

	db := database.MysqlConnect()
	rdb := database.RedisClient()
	courseRepository := impl.NewCourseRepository(db)
	courseService := implservice.NewCourseService(courseRepository, rdb)
	courseController := implcontroller.NewCourseController(courseService)

	api := router.Group("/api")
	api.Post("/courses", courseController.Create)
	api.Get("/courses", courseController.FindAll)

	return router
}
