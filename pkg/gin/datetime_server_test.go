package dateTimeGin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/datetime", GinDateTimeHandler)

	req, err := http.NewRequest("GET", "/datetime", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedDateTime := time.Now()

	var response DateTimeResponse
	assert.NoError(t, json.NewDecoder(rr.Body).Decode(&response), "could not decode response body")

	responseDateTimeStr := response.Date + " " + response.Time

	responseDateTime, err := time.ParseInLocation("02-01-2006 15:04:05", responseDateTimeStr, time.Local)

	assert.NoError(t, err)

	assert.True(t, responseDateTime.After(expectedDateTime.Add(-5*time.Second)) && responseDateTime.Before(expectedDateTime.Add(5*time.Second)))
}
