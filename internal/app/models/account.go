package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number   string `json:"number"`
	Balance  float64 `json:"balance"`
	Owner    string `json:"owner"`
}