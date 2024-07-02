package dateTime

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

	expectedDate := time.Now().Format("02-01-2006")
	expectedTime := time.Now().Format("15:04:05")

	var response DateTimeResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Date != expectedDate {
		t.Errorf("handler returned unexpected date: got %v want %v", response.Date, expectedDate)
	}
	if response.Time[:5] != expectedTime[:5] {
		t.Errorf("handler returned unexpected time: got %v want %v", response.Time, expectedTime)
	}
}
