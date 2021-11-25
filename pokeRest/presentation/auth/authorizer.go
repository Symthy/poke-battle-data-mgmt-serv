package auth

import (
	"net/http"
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	UID  uint   `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

type Authorizer struct {
	service service.UserReadService
}

// Login
func (a Authorizer) SignIn(c echo.Context) error {
	u := new(server.AuthUser)
	if err := c.Bind(u); err != nil {
		return err
	}

	command := command.NewGetUserCommand(*u.Name).SetFilterRequiredFields()
	user, err := a.service.GetUser(*command)
	// Todo: error process
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user.Id() == 0 {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(*u.Password)); err != nil {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		user.Id(),
		user.Name().Value(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	// Todo: error process
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
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

	getCommand := command.NewGetUserCommand(*signupUser.Name).SetFilterRequiredFields()

	u, err := a.service.GetUser(*getCommand)
	// Todo: error process
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if u.Id() != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	createCommand := command.NewCreateUserCommand(
		uint(*signupUser.Id),
		*signupUser.Name,
		*signupUser.Password,
		value.User,
	)
	user, err := a.service.CreateUser(createCommand)
	// Todo: error process
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (a Authorizer) validateUserIdInToken(c echo.Context) error {
	accessUser := c.Get("user").(*jwt.Token)
	claims := accessUser.Claims.(*jwtCustomClaims)
	var uid uint = uint(claims.UID)
	user, err := a.service.GetUserById(uid)
	if user.Id() == 0 {
		return echo.ErrNotFound
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return nil
}
