package dateTime

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GinDateTimeHandler(c *gin.Context) {
	currentTime := time.Now()
	response := DateTimeResponse{
		Date: currentTime.Format("02-01-2006"),
		Time: currentTime.Format("15:04:05"),
	}

	c.JSON(200, response)
}

func StartGinServer() error {
	r := gin.Default()
	r.GET("/datetime", GinDateTimeHandler)
	return r.Run(":8080")
}
