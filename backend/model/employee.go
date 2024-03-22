package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeId uint   `json:"employee_id"`
	Name       string `json:"name"`
}
