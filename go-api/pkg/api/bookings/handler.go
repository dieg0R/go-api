package bookings

import (
	"net/http"
	"strconv"
	"go-api/pkg/models"
	"go-api/pkg/mockDatabase"
	"github.com/gin-gonic/gin"
)

/**
 * @brief GetBookings returns a list of all bookings.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func GetBookings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Bookings)
}

/**
 * @brief PostBookings creates a new booking.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func PostBookings(c *gin.Context) {

	var newBooking models.CreateBooking

	if err := c.ShouldBindJSON(&newBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Booking"})
		return
	}

	if err := models.BookingValidate.Struct(newBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Booking"})
		return
	}

	class, _ := database.FindItemByID(database.Classes, newBooking.ClassId)
	if class == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	classStartDate := class.(*models.Class).StartDate
	classEndDate := class.(*models.Class).EndDate
	bookingDate := newBooking.Date

	if bookingDate.Before(classStartDate) || bookingDate.After(classEndDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Booking date is not within class date range"})
		return
	}

	var booking models.Booking = database.CreateBooking(newBooking)

	database.Bookings = append(database.Bookings, booking)
	c.IndentedJSON(http.StatusCreated, booking)
}

/**
 * @brief GetBookingByID returns a booking by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func GetBookingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingBooking, index := database.FindItemByID(database.Bookings, id)
	if existingBooking == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, database.Bookings[index])
}

/**
 * @brief UpdateBooking updates a booking by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func UpdateBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBooking models.UpdateBooking

	if err := c.ShouldBindJSON(&updatedBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Booking"})
		return
	}

	if err := models.BookingValidate.Struct(updatedBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Booking"})
		return
	}

	existingBooking, index := database.FindItemByID(database.Bookings, id)
	if existingBooking == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	class, _ := database.FindItemByID(database.Classes, updatedBooking.ClassId)
	if class == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	classStartDate := class.(*models.Class).StartDate
	classEndDate := class.(*models.Class).EndDate
	bookingDate := updatedBooking.Date

	if bookingDate.Before(classStartDate) || bookingDate.After(classEndDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Booking date is not within class date range"})
		return
	}

	newBooking:= database.UpdateBooking(updatedBooking, id)
	database.Bookings[index] = newBooking
	c.IndentedJSON(http.StatusOK, newBooking)
}

/**
 * @brief DeleteBooking deletes a booking by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func DeleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingBooking, index := database.FindItemByID(database.Bookings, id)
	if existingBooking == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	database.Bookings = append(database.Bookings[:index], database.Bookings[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}