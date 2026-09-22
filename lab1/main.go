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

	if month < 1 || month > 12 || day < 1 || day > 31 {
		fmt.Println("Некорректная дата")
		return
	}

	if month == 2 && day > 29 {
		fmt.Println("Некорректная дата: февраль не может иметь более 29 дней")
		return
	}

	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		fmt.Println("Некорректная дата: этот месяц не может иметь более 30 дней")
		return
	}

	if year%400 != 0 && year%100 == 0 && month == 2 && day > 28 {
		fmt.Println("Некорректная дата: в невисокосный год февраль не может иметь более 28 дней")
		return
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	result := dayToNewYear(date)
	fmt.Println("До Нового года осталось:", result, "дней")
}
