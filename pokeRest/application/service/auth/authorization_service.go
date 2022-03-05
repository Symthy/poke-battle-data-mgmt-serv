package auth

import (
	"net/http"

	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	"github.com/Symthy/PokeRest/pokeRest/config"
	m_users "github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Todo: remake

type AuthorizationService struct {
	service    users.UserReadService
	authConfig config.AuthConfig
}

func NewAuthorizationService(service users.UserReadService, authConfig config.AuthConfig) AuthorizationService {
	return AuthorizationService{
		service:    service,
		authConfig: authConfig,
	}
}

func (as *AuthorizationService) GenerateToken(name string, password string) (*string, error) {

	command := command.NewGetUserCommand(name).SetFilterRequiredFields()
	user, err := as.service.GetUser(*command)
	if err != nil {
		return nil, err
	}

	if user.Id().Value() == 0 {
		return nil, errs.ThrowServerError(errs.ErrUserNotFound)
	}

	if err := user.ValidatePassword(password); err != nil {
		return nil, errs.ThrowServerError(errs.ErrAuthentication)
	}

	claims := config.NewJwtCustomClaims(user.Id().Value(), user.Name().Value())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(as.authConfig)
	return &t, err
}

func (a AuthorizationService) CreateSignUpUser(name string, password string) (*m_users.User, error) {
	getCommand := command.NewGetUserCommand(name).SetFilterRequiredFields()
	u, err := a.service.GetUser(*getCommand)
	// Todo: error process
	if err != nil {
		return nil, err
	}
	if u.Id().Value() != 0 {
		return nil, &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	createCommand := command.NewCreateUserCommand(
		name,
		password,
		value.User,
	)
	user, err := a.service.SaveUser(createCommand)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a AuthorizationService) ValidateUserIdInToken(c echo.Context) error {
	accessUser := c.Get("user").(*jwt.Token)
	claims := accessUser.Claims.(*config.JwtCustomClaims)
	var uid uint = uint(claims.ID)
	user, err := a.service.GetUserById(uid)
	if user.Id().Value() == 0 {
		return echo.ErrNotFound
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return nil
}
