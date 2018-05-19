package domain

import "time"

//Calendar : ...
type Calendar struct {
	title       string
	visibility  int
	currentTime time.Time
}

//NewCalendar : constructor for Calendar
func NewCalendar() *Calendar {
	return &Calendar{}
}
