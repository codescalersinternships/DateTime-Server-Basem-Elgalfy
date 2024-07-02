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

	expectedDate := time.Now().Format("02-01-2006")
	expectedTime := time.Now().Format("15:04:05")

	var response DateTimeResponse
	assert.NoError(t, json.NewDecoder(rr.Body).Decode(&response), "could not decode response body")

	assert.Equal(t, expectedDate, response.Date)
	assert.Equal(t, expectedTime[:5], response.Time[:5])
}
