package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

type Job struct {
	gorm.Model
	Title   	string
	Description string
	MinSalaray  int
	MaxSalaray  int
}

type Employee struct {
	gorm.Model
	Name   string
	DOB    time.Time
	salary int
	JobID  int
	Job    Job `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type LoginHistory struct {
	gorm.Model
	Description string
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type JobHistory struct {
	gorm.Model
	Description string
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
