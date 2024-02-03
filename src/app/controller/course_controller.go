package controller

import "github.com/gofiber/fiber/v2"

type CourseController interface {
	Create(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
