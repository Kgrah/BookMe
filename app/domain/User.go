package domain

//User : ...
type User struct {
	username      string
	password      []byte
	groups        []Group
	friends       Graph
	bookedEvents  []Event
	createdEvents []Event
	services      []*Service
}

//NewUser : constructor for User
func NewUser() *User {
	return &User{}
}

func (user *User) book(event Event) {
	event.bookers = append(event.bookers, user)
}
