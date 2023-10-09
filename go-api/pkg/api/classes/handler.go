package classes

import (
	"net/http"
	"strconv"
	"go-api/pkg/models"
	"go-api/pkg/mockDatabase"
	"github.com/gin-gonic/gin"
)

/**
 * @brief GetClasses returns a list of all classes.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func GetClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Classes)
}

/**
 * @brief PostClasses creates a new class.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func PostClasses(c *gin.Context) {
	var newClass models.CreateClass

	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Class"})
		return
	}

	if err := models.ClassValidate.Struct(newClass); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Class"})
		return
	}

	if newClass.StartDate.After(newClass.EndDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "StartDate must be before EndDate"})
		return
	}

	var class models.Class = database.CreateClass(newClass)

	database.Classes = append(database.Classes, class)
	c.IndentedJSON(http.StatusCreated, class)
}

/**
 * @brief GetClassesByID returns a class by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func GetClassesByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingClass, index := database.FindItemByID(database.Classes, id)
	if existingClass == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, database.Classes[index])
}

/**
 * @brief UpdateClass updates a class by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func UpdateClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedClass models.UpdateClass

	if err := c.ShouldBindJSON(&updatedClass); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Class"})
		return
	}

	if err := models.ClassValidate.Struct(updatedClass); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Class"})
		return
	}

	existingClass, index := database.FindItemByID(database.Classes, id)
	if existingClass == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	if updatedClass.StartDate.After(updatedClass.EndDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "StartDate must be before EndDate"})
		return
	}

	newclass := database.UpdateClass(updatedClass, id)
	database.Classes[index] = newclass
	c.IndentedJSON(http.StatusOK, newclass)
}

/**
 * @brief DeleteClass deletes a class by its ID.
 * 
 * @param c *gin.Context: The Gin HTTP context.
 */
func DeleteClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingClass, index := database.FindItemByID(database.Classes, id)
	if existingClass == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	database.Classes = append(database.Classes[:index], database.Classes[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Class deleted"})
}