package main

import (
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
