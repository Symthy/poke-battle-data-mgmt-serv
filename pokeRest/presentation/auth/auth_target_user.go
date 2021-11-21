package auth

type AuthTargetUser struct {
	id       uint
	name     string
	password string
}

func NewAuthTargetUser(id uint, name string, password string) AuthTargetUser {
	return AuthTargetUser{
		id:       id,
		name:     name,
		password: password,
	}
}

func (u AuthTargetUser) Id() uint {
	return u.id
}

func (u AuthTargetUser) Name() string {
	return u.name
}

func (u AuthTargetUser) Password() string {
	return u.password
}
