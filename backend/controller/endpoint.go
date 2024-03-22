package controller

import (
	"employees/model"
	"employees/service"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Employee struct {
	employeeService *service.Employee
}

func NewEmployeeController(employeeService *service.Employee) *Employee {
	return &Employee{
		employeeService: employeeService,
	}
}

func (e *Employee) Create(ctx *fiber.Ctx) error {
	var employeeDetails model.Employee
	err := json.Unmarshal(ctx.Body(), &employeeDetails)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(nil)
	}
	err = e.employeeService.Create(employeeDetails)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(nil)
	}
	return ctx.Status(http.StatusCreated).JSON(nil)
}

func (e *Employee) GetAll(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(e.employeeService.GetAll())
}
