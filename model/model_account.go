package model

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Name string `json:"name"`
	Balance float64 `json:"balance"`
}