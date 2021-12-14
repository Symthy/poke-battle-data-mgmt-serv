package value

import (
	"fmt"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/stretchr/testify/suite"
)

type UserNameTestSuite struct {
	suite.Suite
}

func TestUserNameTestSuite(t *testing.T) {
	suite.Run(t, new(UserNameTestSuite))
}

func (suite *UserNameTestSuite) TestNewUserName() {
	suite.Run("return UserName", func() {
		expected := "dummy_user"
		name, err := value.NewUserName(expected)

		if err != nil {
			suite.Fail(err.Error())
		}
		if expected != (*name).Value() {
			suite.Fail("unmatched name")
			fmt.Printf("expected:%v\n", expected)
			fmt.Printf("actual:  %v\n", (*name).Value())
		}
	})

	suite.Run("return error", func() {
		expected := "bad@name."
		name, err := value.NewUserName(expected)

		if err == nil {
			suite.Fail("no error")
			fmt.Printf("actual:  %v\n", (*name).Value())
		}
	})
}
