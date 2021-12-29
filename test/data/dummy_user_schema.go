package data

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/collection"
)

func DummyUser1(filterFields ...string) schema.User {
	var id uint = 0
	if len(filterFields) == 0 || collection.ListContains(filterFields, "id") {
		id = 1
	}
	var displayName *string = nil
	if len(filterFields) == 0 || collection.ListContains(filterFields, "displayName") {
		displayName = new(string)
		*displayName = "dummy dummy user"
	}
	var email *string = nil
	if len(filterFields) == 0 || collection.ListContains(filterFields, "email") {
		email = new(string)
		*email = "test@test.com"
	}
	var profile *string = nil
	if len(filterFields) == 0 || collection.ListContains(filterFields, "profile") {
		profile = new(string)
		*profile = "detail\nprofile"
	}

	dummyUser := schema.User{
		Name:        "dummy_user",
		DisplayName: displayName,
		Email:       email,
		Profile:     profile,
		Role:        enum.User,
	}
	dummyUser.ID = id
	return dummyUser
}
