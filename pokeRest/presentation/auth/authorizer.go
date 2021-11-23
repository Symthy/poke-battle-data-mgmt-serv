package auth

import (
	"net/http"
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
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

type User struct {
	ID       uint   `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
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
	// Todo: error process
	user, _ := a.service.GetUser(*command)
	if user.Id() == 0 {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(*u.Password))
	if err != nil {
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
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// first register
func (a Authorizer) SignUp(c echo.Context) error {
	signupUser := new(User)
	if err := c.Bind(signupUser); err != nil {
		return err
	}

	if signupUser.Name == "" || signupUser.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	command := command.NewGetUserCommand(signupUser.Name).SetFilterRequiredFields()
	// Todo: error process
	if u, _ := a.service.GetUser(*command); u.Id() != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	// Todo: command
	user, err := a.service.CreateUser(signupUser)

	return c.JSON(http.StatusCreated, user)
}

func (a Authorizer) validateUserIdInToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	var uid uint = uint(claims.UID)
	if user := a.controller.FindUser(&User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
	return nil
}
