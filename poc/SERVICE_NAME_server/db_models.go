package main

import (
	"time"
)

type Model1 struct {
	Field1 uint   `gorm:"primary_key"`
	Field2 string `sql:"size:11"`
	Field3 time.Time
}

type Model2 struct {
	Field4 uint   `gorm:"primary_key"`
	Field5 string `sql:"size:11"`
	Field6 time.Time
}
