package models

import (
	"time"
    "github.com/go-playground/validator/v10"
)

var ClassValidate *validator.Validate = validator.New()

type Class struct {
	ID         int `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required,alphanum,max=20"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
	Capacity   int    `json:"capacity" validate:"required"`
}

type CreateClass struct {
	Name       string `json:"name" validate:"required,alphanum,max=20"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
	Capacity   int    `json:"capacity" validate:"required"`
}

type UpdateClass struct {
	Name       string `json:"name" validate:"required,alphanum,max=20"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
	Capacity   int    `json:"capacity" validate:"required"`
}