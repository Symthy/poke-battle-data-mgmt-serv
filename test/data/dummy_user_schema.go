package data

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
)

func DummyUser1() schema.User {
	displayName := "dummy dummy user"
	email := "test@test.com"
	profile := "detail\nprofile"
	dummyUser := schema.User{
		Name:        "dummy_user",
		DisplayName: &displayName,
		Email:       &email,
		Profile:     &profile,
		Role:        enum.User,
	}
	dummyUser.ID = 1
	return dummyUser
}
