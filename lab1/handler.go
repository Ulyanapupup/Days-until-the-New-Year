package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func daysHandler(w http.ResponseWriter, r *http.Request) {
	dateString := r.URL.Query().Get("date")
	if dateString == "" {
		dateString = time.Now().Format("2006-01-02")
	}

	date, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := struct {
			Error string `json:"error"`
		}{
			Error: "invalid date",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	days := dayToNewYear(date)
	response := struct {
		Date string `json:"date"`
		Days int    `json:"days"`
	}{
		Date: dateString,
		Days: days,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	json.NewEncoder(w).Encode(response)
}
