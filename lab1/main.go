package main

import (
	"net/http"
	"time"
)

func dayToNewYear(date time.Time) int {
	newYear := time.Date(date.Year()+1, 1, 1, 0, 0, 0, 0, date.Location())
	result := newYear.Sub(date).Hours() / 24
	return int(result)
}

func main() {
	http.HandleFunc("/days", daysHandler)
	http.ListenAndServe(":8080", nil)
}
