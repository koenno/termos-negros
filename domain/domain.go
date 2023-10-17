package domain

import "time"

type Menu []DayMenu

type DayMenu struct {
	Date  time.Time
	Meals []Meal
}

type Meal struct {
	Name        string
	Ingredients string
}
