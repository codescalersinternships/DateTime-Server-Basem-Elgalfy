package dateTime

import (
	"encoding/json"
	"net/http"
	"time"
)

type DateTimeResponse struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func DateTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	response := DateTimeResponse{
		Date: currentTime.Format("02-01-2006"),
		Time: currentTime.Format("15:04:05"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StartServer() error {
	http.HandleFunc("/datetime", DateTimeHandler)
	return http.ListenAndServe(":8080", nil)
}
