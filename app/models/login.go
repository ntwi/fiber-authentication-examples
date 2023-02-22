package models

import (
	"github.com/go-playground/validator/v10"
)

type LoginModel struct {
	Email    string `json:"email" gorm:"unique;" validate:"required,email,min=6,max=32"`
	Password string `json:"password" gorm:"type:text;" validate:"required,min=6"`
}

func (l LoginModel) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
