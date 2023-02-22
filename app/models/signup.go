package models

import (
	"github.com/go-playground/validator/v10"
)

type SignUpModel struct {
	Username         string `json:"username" validate:"required,min=3,max=32"`
	Email            string `json:"email" validate:"required,email,min=6,max=32"`
	Password         string `json:"password" validate:"required,min=6"`
	ReferralUsername string `json:"referralUsername" validate:"omitempty,min=3,max=32"`
}

func (l SignUpModel) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
