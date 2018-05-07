package domain

import "time"

//Event : ...
type Event struct {
	title       string
	occurances  [][]time.Time
	bookers     []*User
	invitees    []User
	creator     User
	restriction int
	visibility  int
}

//NewEvent : the constructor for Event
func NewEvent() *Event {
	return &Event{}
}
