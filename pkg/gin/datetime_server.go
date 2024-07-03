// package dateTimeGin is a server that returns dateTime in the standard JSON when requested by /dateTime
package dateTimeGin

import (
	"time"

	"github.com/gin-gonic/gin"
)

// DateTimeResponse is the response struct for the DateTimeHandler
type DateTimeResponse struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// GinDateTimeHandler is the handler for the /datetime endpoint
func GinDateTimeHandler(c *gin.Context) {
	currentTime := time.Now()
	response := DateTimeResponse{
		Date: currentTime.Format("02-01-2006"),
		Time: currentTime.Format("15:04:05"),
	}

	c.JSON(200, response)
}

// StartGinServer starts the server
func StartGinServer() error {
	r := gin.Default()
	r.GET("/datetime", GinDateTimeHandler)
	return r.Run(":8080")
}
