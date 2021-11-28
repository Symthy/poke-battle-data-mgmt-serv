package auth

import (
	"net/http"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/auth"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
	"github.com/labstack/echo/v4"
)

type Authorizer struct {
	authService auth.AuthorizationService
}

func NewAuthorizer(service auth.AuthorizationService) Authorizer {
	return Authorizer{authService: service}
}

// Login
func (a Authorizer) SignIn(c echo.Context) error {
	u := new(server.AuthUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	token, err := a.authService.GenerateToken(*u.Name, *u.Password)

	// Todo: error process
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// first register
func (a Authorizer) SignUp(c echo.Context) error {
	signupUser := new(server.AuthUser)
	if err := c.Bind(signupUser); err != nil {
		return err
	}

	if *signupUser.Name == "" || *signupUser.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	user, err := a.authService.CreateSignUpUser(*signupUser.Name, *signupUser.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, presentation.ConvertUserToResponse(*user))
}

// Todo
// func errorHandling(err error) error {
// }
