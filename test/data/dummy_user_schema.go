package data

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
)

func DummyUser1(filterFields ...string) *schema.User {
	var id uint64 = 0
	if len(filterFields) == 0 || lists.Contains(filterFields, "id") {
		id = 1
	}
	var displayName *string = nil
	if len(filterFields) == 0 || lists.Contains(filterFields, "displayName") {
		displayName = new(string)
		*displayName = "dummy dummy user"
	}
	var email *string = nil
	if len(filterFields) == 0 || lists.Contains(filterFields, "email") {
		email = new(string)
		*email = "test@test.com"
	}
	var profile *string = nil
	if len(filterFields) == 0 || lists.Contains(filterFields, "profile") {
		profile = new(string)
		*profile = "detail\nprofile"
	}

	dummyUser := &schema.User{
		UserSchema: schema.UserSchema{
			Name:        "dummy_user",
			DisplayName: displayName,
			Email:       email,
			Profile:     profile,
			Role:        enum.User,
		},
	}
	dummyUser.ID = id
	return dummyUser
}
