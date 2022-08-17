package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ users.IUserNotification = (*UserSchemaBuilder)(nil)

type UserSchemaBuilder struct {
	*schema.User
}

func NewUserSchemaBuilder() *UserSchemaBuilder {
	return &UserSchemaBuilder{&schema.User{}}
}

func (b UserSchemaBuilder) SetId(id identifier.UserId) {
	b.ID = id.Value()
}
func (b UserSchemaBuilder) SetName(name value.UserName) {
	b.Name = name.Value()
}
func (b UserSchemaBuilder) SetDisplayName(displayName *string) {
	b.DisplayName = displayName
}
func (b UserSchemaBuilder) SetProfile(profile *string) {
	b.Profile = profile
}
func (b UserSchemaBuilder) SetRole(role value.Role) {
	b.Role = enum.Role(role.String())
}

func (b UserSchemaBuilder) Build() *schema.User {
	return b.User
}
