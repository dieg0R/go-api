package database

import (
	"time"
	"go-api/pkg/models"
)

var bookingIDCounter = 3
var classIDCounter = 3

var Bookings = []models.Booking{
    {ID: 1, Name: "Diego", ClassId: 1, Date: time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC)},
    {ID: 2, Name: "Martin", ClassId: 2, Date: time.Date(2023, 10, 7, 20, 0, 0, 0, time.UTC)},
    {ID: 3, Name: "Joaquin", ClassId: 3, Date: time.Date(2023, 10, 11, 11, 0, 0, 0, time.UTC)},
}

var Classes = []models.Class{
	{ID: 1, Name: "Yoga", StartDate: time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC), Capacity: 10},
	{ID: 2, Name: "Pilates", StartDate: time.Date(2023, 10, 7, 20, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 10, 21, 0, 0, 0, time.UTC), Capacity: 8},
	{ID: 3, Name: "Boxing", StartDate: time.Date(2023, 10, 11, 11, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 12, 12, 0, 0, 0, time.UTC), Capacity: 12},
}

func FindItemByID(list interface{}, id int) (interface{}, int) {
	switch items := list.(type) {
	case []models.Booking:
		for index, item := range items {
			if item.ID == id {
				return &item, index
			}
		}
	case []models.Class:
		for index, item := range items {
			if item.ID == id {
				return &item, index
			}
		}
	}
	return nil, -1
}

func CreateBooking(newBooking models.CreateBooking) models.Booking {
	bookingIDCounter++
    booking := models.Booking{
        ID:        bookingIDCounter,
        Name:      newBooking.Name,
        ClassId:   newBooking.ClassId,
        Date:      newBooking.Date,
    }
    return booking
}

func UpdateBooking(newBooking models.UpdateBooking, index int) models.Booking {
    booking := models.Booking{
		ID:        		index,
		Name:      		newBooking.Name,
        ClassId:   		newBooking.ClassId,
        Date:      		newBooking.Date,
    }
    return booking
}

func CreateClass(newClass models.CreateClass) models.Class {
	classIDCounter++
    class := models.Class{
		ID:        		classIDCounter,
		Name:      		newClass.Name,
        StartDate:      newClass.StartDate,
        EndDate:   		newClass.EndDate,
        Capacity:      	newClass.Capacity,
    }
    return class
}

func UpdateClass(newClass models.UpdateClass, index int) models.Class {
    class := models.Class{
		ID:        		index,
		Name:      		newClass.Name,
        StartDate:      newClass.StartDate,
        EndDate:   		newClass.EndDate,
        Capacity:      	newClass.Capacity,
    }
    return class
}