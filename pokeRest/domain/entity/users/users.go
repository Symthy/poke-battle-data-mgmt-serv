package users

type Users struct {
	items []User
}

func NewUsers(items []User) Users {
	return Users{items: items}
}

func (u Users) Items() []User {
	return u.items
}

func (u Users) First() User {
	return u.items[0]
}
