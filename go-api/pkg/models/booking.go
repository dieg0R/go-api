package models

import (
	"time"
    "github.com/go-playground/validator/v10"
)

var BookingValidate *validator.Validate = validator.New()

type Booking struct {
	ID       	int `json:"id" validate:"required"`
	Name      	string `json:"name" validate:"required,alphanum,max=20"`
	ClassId 	int `json:"class_id" validate:"required"`
	Date      	time.Time `json:"date" validate:"required"`
}

type CreateBooking struct {
	Name      	string `json:"name" validate:"required,alphanum,max=20"`
	ClassId 	int `json:"class_id" validate:"required"`
	Date      	time.Time `json:"date" validate:"required"`
}

type UpdateBooking struct {
	Name      	string `json:"name" validate:"required,alphanum,max=20"`
	ClassId 	int `json:"class_id" validate:"required"`
	Date      	time.Time `json:"date" validate:"required"`
}