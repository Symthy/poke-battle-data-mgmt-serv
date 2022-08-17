package users

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type IUserNotification interface {
	SetId(identifier.UserId)
	SetName(value.UserName)
	SetDisplayName(*string)
	SetProfile(*string)
	SetRole(value.Role)
}
