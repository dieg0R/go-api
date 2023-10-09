package api

import (
	"go-api/pkg/api/bookings"
	"go-api/pkg/api/classes"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/classes", classes.GetClasses)
		api.GET("/classes/:id", classes.GetClassesByID)
		api.POST("/classes", classes.PostClasses)
		api.PUT("/classes/:id", classes.UpdateClass)
		api.DELETE("/classes/:id", classes.DeleteClass)

		api.GET("/bookings", bookings.GetBookings)
		api.GET("/bookings/:id", bookings.GetBookingByID)
		api.POST("/bookings", bookings.PostBookings)
		api.PUT("/bookings/:id", bookings.UpdateBooking)
		api.DELETE("/bookings/:id", bookings.DeleteBooking)
	}

	return router
}