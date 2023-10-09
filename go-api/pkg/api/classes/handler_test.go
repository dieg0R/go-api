package classes

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

func TestGetClassesByID(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.GET("/classes/:id", GetClassesByID)

	// Create a GET request to retrieve booking with ID 1
	req, _ := http.NewRequest(http.MethodGet, "/classes/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into the expected structure
	var response models.Class
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Yoga", response.Name)
	assert.Equal(t, time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC), response.StartDate)
	assert.Equal(t, time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC), response.EndDate)
	assert.Equal(t, 10, response.Capacity)
}

func TestGetClasses(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.GET("/classes", GetClasses)

	// Create a GET request to retrieve all classes
	req, _ := http.NewRequest(http.MethodGet, "/classes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into a slice of bookings
	var response []models.Class
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions to compare the response with the database.Bookings
	assert.Equal(t, len(database.Classes), len(response))
	for i, class := range database.Classes {
		assert.Equal(t, class.ID, response[i].ID)
		assert.Equal(t, class.Name, response[i].Name)
		assert.Equal(t, class.StartDate, response[i].StartDate)
		assert.Equal(t, class.EndDate, response[i].EndDate)
		assert.Equal(t, class.Capacity, response[i].Capacity)
	}
}

func TestPostClasses(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.POST("/classes", PostClasses)

	// Define a new class  for testing
	var newClass = models.CreateClass{
		Name:    "TestClass",
		StartDate: time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC),
		Capacity: 8,
	}

	newClassJSON, _ := json.Marshal(newClass)

	// Create a POST request to create a new class
	req, _ := http.NewRequest(http.MethodPost, "/classes", bytes.NewReader(newClassJSON))
	
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Created (201)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the response body into the created booking
	var createdClass models.Class
	responseBody := w.Body.Bytes()

	err := json.Unmarshal(responseBody, &createdClass)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the created booking
	assert.NotNil(t, createdClass)
	assert.NotZero(t, createdClass.ID)
	assert.Equal(t, newClass.Name, createdClass.Name)
	assert.Equal(t, newClass.StartDate, createdClass.StartDate)
	assert.Equal(t, newClass.EndDate, createdClass.EndDate)
	assert.Equal(t, newClass.Capacity, createdClass.Capacity)
}

func TestUpdateClass(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.PUT("/classes/:id", UpdateClass)

	// Define the updated class  for testing
	var updatedClass = models.UpdateClass{
		Name:    "TestClassUpdated",
		StartDate: time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC),
		Capacity: 10,
	}

	// Convert the updated class to JSON
	updatedClassJSON, _ := json.Marshal(updatedClass)

	// Create a PUT request to update an existing class
	req, _ := http.NewRequest(http.MethodPut, "/classes/1", bytes.NewBuffer(updatedClassJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body into the updated class
	var updated models.Class
	responseBody := w.Body.Bytes()
	err := json.Unmarshal(responseBody, &updated)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the updated class
	assert.NotNil(t, updated)
	assert.NotZero(t, updated.ID)
	assert.Equal(t, updatedClass.Name, updated.Name)
	assert.Equal(t, updatedClass.StartDate, updated.StartDate)
	assert.Equal(t, updatedClass.EndDate, updated.EndDate)
	assert.Equal(t, updatedClass.Capacity, updated.Capacity)
}

func TestDeleteClass(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.DELETE("/classes/:id", DeleteClass)

	// Create a DELETE request to delete an existing class
	req, _ := http.NewRequest(http.MethodDelete, "/classes/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostClassInvalidDateOrder(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.POST("/classes", PostClasses)

	// Define a new class with StartDate after EndDate for testing
	var newClass = models.CreateClass{
		Name:      "TestClassInvalidDate",
		StartDate: time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
		Capacity:  8,
	}

	newClassJSON, _ := json.Marshal(newClass)

	// Create a POST request to create a new class with invalid date order
	req, _ := http.NewRequest(http.MethodPost, "/classes", bytes.NewReader(newClassJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Bad Request (400)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateClassInvalidDateOrde(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()
	router.PUT("/classes/:id", UpdateClass)

	// Define a new class with StartDate after EndDate for testing
	var updatedClass = models.CreateClass{
		Name:      "TestClassInvalidDate",
		StartDate: time.Date(2023, 10, 16, 17, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 10, 6, 16, 0, 0, 0, time.UTC),
		Capacity:  8,
	}

	// Convert the updated class to JSON
	updatedClassJSON, _ := json.Marshal(updatedClass)

	// Create a PUT request to update an existing class with invalid date order
	req, _ := http.NewRequest(http.MethodPut, "/classes/1", bytes.NewBuffer(updatedClassJSON))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is Bad Request (404)
	assert.Equal(t, http.StatusNotFound, w.Code)
}