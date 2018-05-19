package domain

//Group : ...
type Group struct {
	members          []User
	tile             string
	restrictionLevel int
	invitees         []*User
	User
}

//NewGroup : constructor for Group
func NewGroup() *Group {
	return &Group{}
}

func (group *Group) addUser(user *User) {
	group.invitees = append(group.invitees, user)
}
