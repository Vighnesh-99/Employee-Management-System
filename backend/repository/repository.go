package repository

import (
	"employees/model"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Employee struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *Employee {
	return &Employee{
		db: db,
	}
}

func (e Employee) Save(employee model.Employee) error {
	tx := e.db.Create(&employee)
	if tx.Error != nil {
		log.Println("error saving employees in db ", tx.Error.Error())
		return tx.Error
	}
	return nil
}

func (e Employee) FetchAll() []model.Employee {
	var employees []model.Employee
	tx := e.db.Find(&employees)
	if tx.Error != nil {
		log.Println("error fetching employees ", tx.Error.Error())
		return nil
	}
	return employees
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.Employee{})
	if err != nil {
		log.Fatalln(fmt.Sprintf("error running employee migrations %s", err.Error()))
	}
}
