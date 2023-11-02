package models

import (
	"gorm.io/gorm"
)

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
