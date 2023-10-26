package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	PhoneNumber  string `json:"phonenumber"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validste:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}
type Job struct {
	gorm.Model
	JobId     uint `gorm:"primaryKey:autoIncrement"`
	JobName   string
	CompanyID uint
}
type NewJob struct {
	JobName string `json:"JobName" validate:"required"`
}
type Company struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Location string
	Jobs     []Job
}
type NewCompany struct {
	Name     string `gorm:"unique"`
	Location string
	Jobs     []Job
}
