package impl

import (
	"Redis/redis-impl-go/src/app/controller"
	"Redis/redis-impl-go/src/app/model/web"
	"Redis/redis-impl-go/src/app/service"
	"Redis/redis-impl-go/src/app/utils/errors"
	"Redis/redis-impl-go/src/app/utils/response"
	"github.com/gofiber/fiber/v2"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) Create(ctx *fiber.Ctx) error {
	input := web.CourseCreateInput{}
	err := ctx.BodyParser(&input)
	errors.PanicIfError(err)

	courseResponse := c.CourseService.Create(input)

	ctx.JSON(response.APIResponse(200, "course has successfully created", &courseResponse))
	return nil
}

func (c *CourseControllerImpl) FindAll(ctx *fiber.Ctx) error {
	courseResponses := c.CourseService.FindAll()

	ctx.JSON(response.APIResponse(200, "list of courses", &courseResponses))
	return nil
}

func NewCourseController(courseService service.CourseService) controller.CourseController {
	return &CourseControllerImpl{CourseService: courseService}
}
