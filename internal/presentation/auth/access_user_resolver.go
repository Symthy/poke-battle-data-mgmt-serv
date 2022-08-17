package auth

import (
	"github.com/Symthy/PokeRest/internal/application/service/users"
	"github.com/Symthy/PokeRest/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AccessUserResolver struct {
	userService users.UserReadService
}

func NewAccessUserResolver(userService users.UserReadService) AccessUserResolver {
	return AccessUserResolver{userService}
}

func (r AccessUserResolver) ResolveId(c echo.Context) (uint64, error) {
	accessUser := c.Get("user").(*jwt.Token)
	claims := accessUser.Claims.(*config.JwtCustomClaims)
	var uid uint64 = uint64(claims.ID)
	user, err := r.userService.GetUserById(uid)
	if err != nil {
		return 0, err
	}
	id := user.Id().Value()
	return id, nil
}
