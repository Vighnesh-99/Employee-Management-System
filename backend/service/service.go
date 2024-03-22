package service

import (
	"employees/model"
	"employees/repository"
)

type Employee struct {
	employeeRepository *repository.Employee
}

func NewEmployeeService(employeeRepository *repository.Employee) *Employee {
	return &Employee{
		employeeRepository: employeeRepository,
	}
}

func (e Employee) Create(employee model.Employee) error {
	return e.employeeRepository.Save(employee)
}

func (e Employee) GetAll() []model.Employee {
	return e.employeeRepository.FetchAll()
}
