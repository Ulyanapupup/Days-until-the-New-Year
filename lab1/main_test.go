package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDayToNewYear(t *testing.T) {
	cases := []struct {
		name string
		date time.Time
		want int
	}{
		{"конец года", time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local), 1},
		{"начало года", time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local), 365},
		{"високосный год", time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local), 366},
		{"до 29 февраля", time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local), 335},
		{"после 29 февраля", time.Date(2024, 3, 1, 0, 0, 0, 0, time.Local), 306},
		{"обычный день", time.Date(2026, 6, 15, 0, 0, 0, 0, time.Local), 200},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := dayToNewYear(tc.date)
			if got != tc.want {
				t.Errorf("для %v получили %d, должно быть %d",
					tc.date.Format("02.01.2006"), got, tc.want)
			}
		})
	}
}

func TestDaysHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/days?date=2026-01-01", nil)
	recorder := httptest.NewRecorder()

	daysHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("получен статус %d, ожидался %d",
			recorder.Code, http.StatusOK)
	}

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("получен Content-Type %q, ожидался application/json", contentType)
	}

	var response struct {
		Date string `json:"date"`
		Days int    `json:"days"`
	}

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("ошибка разбора JSON: %v", err)
	}

	if response.Date != "2026-01-01" {
		t.Errorf("получена дата %s, ожидалась 2026-01-01", response.Date)
	}

	if response.Days != 365 {
		t.Errorf("получено %d дней, ожидалось 365", response.Days)
	}
}

func TestDaysHandlerInvalidDate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/days?date=hello", nil)
	recorder := httptest.NewRecorder()

	daysHandler(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("получен статус %d, ожидался %d",
			recorder.Code, http.StatusBadRequest)
	}

	var response struct {
		Error string `json:"error"`
	}

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("ошибка разбора JSON: %v", err)
	}

	if response.Error != "invalid date" {
		t.Errorf("получена ошибка %q, ожидалась %q",
			response.Error, "invalid date")
	}
}

func TestDaysHandlerWithoutDate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/days", nil)
	recorder := httptest.NewRecorder()

	daysHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("получен статус %d, ожидался %d",
			recorder.Code, http.StatusOK)
	}

	var response struct {
		Date string `json:"date"`
		Days int    `json:"days"`
	}

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("ошибка разбора JSON: %v", err)
	}

	now := time.Now()
	today := now.Format("2006-01-02")

	if response.Date != today {
		t.Errorf("получена дата %s, ожидалась %s",
			response.Date, today)
	}

	expectedDate, err := time.Parse("2006-01-02", today)
	if err != nil {
		t.Fatalf("ошибка разбора текущей даты: %v", err)
	}

	expectedDays := dayToNewYear(expectedDate)

	if response.Days != expectedDays {
		t.Errorf("получено %d дней, ожидалось %d",
			response.Days, expectedDays)
	}
}
