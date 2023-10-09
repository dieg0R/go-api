package bookings

import (
	"bytes"
	"go-api/pkg/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"go-api/pkg/mockDatabase"
	"time"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBookingByID(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.GET("/bookings/:id", GetBookingByID)

	// Create a GET request to retrieve booking with ID 1
	req, _ := http.NewRequest(http.MethodGet, "/bookings/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into the expected structure
	var response models.Booking
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Diego", response.Name)
	assert.Equal(t, 1, response.ClassId)
	assert.Equal(t, time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC), response.Date)
}

func TestGetBookings(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.GET("/bookings", GetBookings)

	// Create a GET request to retrieve all bookings
	req, _ := http.NewRequest(http.MethodGet, "/bookings", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into a slice of bookings
	var response []models.Booking
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions to compare the response with the database.Bookings
	assert.Equal(t, len(database.Bookings), len(response))
	for i, booking := range database.Bookings {
		assert.Equal(t, booking.ID, response[i].ID)
		assert.Equal(t, booking.Name, response[i].Name)
		assert.Equal(t, booking.ClassId, response[i].ClassId)
		assert.Equal(t, booking.Date, response[i].Date)
	}
}

func TestPostBookings(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.POST("/bookings", PostBookings)

	// Define a new booking  for testing
	var newBooking = models.CreateBooking{
		Name:    "TestBooking",
		ClassId: 1,
		Date:    time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
	}
	newBookingJSON, _ := json.Marshal(newBooking)

	// Create a POST request to create a new booking
	req, _ := http.NewRequest(http.MethodPost, "/bookings", bytes.NewReader(newBookingJSON))
	
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Created (201)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the response body into the created booking
	var createdBooking models.Booking
	responseBody := w.Body.Bytes()

	err := json.Unmarshal(responseBody, &createdBooking)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the created booking
	assert.NotNil(t, createdBooking)
	assert.NotZero(t, createdBooking.ID)
	assert.Equal(t, newBooking.Name, createdBooking.Name)
	assert.Equal(t, newBooking.ClassId, createdBooking.ClassId)
	assert.Equal(t, newBooking.Date, createdBooking.Date)
}

func TestUpdateBooking(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Define the updated booking  for testing
	var updatedBooking = models.UpdateBooking{
		Name:    "TestBookingUpdated",
		ClassId: 1,
		Date:    time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
	}

	// Convert the updated booking to JSON
	updatedBookingJSON, _ := json.Marshal(updatedBooking)

	// Create a PUT request to update an existing booking
	req, _ := http.NewRequest(http.MethodPut, "/bookings/1", bytes.NewBuffer(updatedBookingJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into the updated booking
	var updated models.Booking
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &updated)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the updated booking
	assert.NotNil(t, updated)
	assert.Equal(t, updatedBooking.Name, updated.Name)
	assert.Equal(t, updatedBooking.ClassId, updated.ClassId)
	assert.Equal(t, updatedBooking.Date, updated.Date)
}

func TestDeleteBooking(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.DELETE("/bookings/:id", DeleteBooking)

	// Create a DELETE request to delete an existing booking
	req, _ := http.NewRequest(http.MethodDelete, "/bookings/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostBookingsInvalidClassID(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.POST("/bookings", PostBookings)

	// Define a new booking with an invalid ClassId (ClassId 50 does not exist)
	var newBooking = models.CreateBooking{
		Name:    "TestBooking",
		ClassId: 50, // Invalid ClassId
		Date:    time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
	}
	newBookingJSON, _ := json.Marshal(newBooking)

	// Create a POST request to create a new booking with an invalid ClassId
	req, _ := http.NewRequest(http.MethodPost, "/bookings", bytes.NewReader(newBookingJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Not Found Request (404)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateBookingInvalidClassID(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Define an updated booking with an invalid ClassId (ClassId 50 does not exist)
	var updatedBooking = models.UpdateBooking{
		Name:    "TestBookingUpdated",
		ClassId: 50, // Invalid ClassId
		Date:    time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
	}

	// Convert the updated booking to JSON
	updatedBookingJSON, _ := json.Marshal(updatedBooking)

	// Create a PUT request to update an existing booking with an invalid ClassId
	req, _ := http.NewRequest(http.MethodPut, "/bookings/1", bytes.NewBuffer(updatedBookingJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Not Found Request (404)
	assert.Equal(t, http.StatusNotFound, w.Code)
}