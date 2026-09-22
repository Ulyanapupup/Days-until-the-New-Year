package main

import (
	"fmt"
	"time"
)

func dayToNewYear(date time.Time) int {
	newYear := time.Date(date.Year()+1, 1, 1, 0, 0, 0, 0, date.Location())
	result := newYear.Sub(date).Hours() / 24
	return int(result)
}

func main() {
	var day int
	var month int
	var year int

	fmt.Print("Введите дату в формате \"20 9 2026\":  ")
	_, err := fmt.Scan(&day, &month, &year)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	result := dayToNewYear(date)
	fmt.Println("До Нового года осталось:", result, "дней")
}
