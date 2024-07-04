package dateTimeHttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDateTimeHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/datetime", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DateTimeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedDateTime := time.Now()

	var response DateTimeResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	responseDateTimeStr := response.Date + " " + response.Time

	responseDateTime, err := time.ParseInLocation("02-01-2006 15:04:05", responseDateTimeStr, time.Local)
	if err != nil {
		t.Fatalf("could not parse response time: %v", err)
	}

	if responseDateTime.Before(expectedDateTime.Add(-5*time.Second)) || responseDateTime.After(expectedDateTime.Add(5*time.Second)) {
		t.Errorf("response time is not within 5 seconds of the expected time: got %v want %v", responseDateTime, expectedDateTime)
	}
}
