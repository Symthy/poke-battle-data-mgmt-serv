package value

import (
	"fmt"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/stretchr/testify/suite"
)

type EmailTestSuite struct {
	suite.Suite
}

func TestEmailTestSuite(t *testing.T) {
	suite.Run(t, new(EmailTestSuite))
}

func (suite *EmailTestSuite) TestNewEmail() {
	expected := "test@test.com"
	email, err := value.NewEmail(expected)

	if err != nil {
		suite.Fail(err.Error())
	}
	if expected != email.Value() {
		suite.Fail("unmatched email")
		fmt.Printf("expected:%v\n", expected)
		fmt.Printf("actual:  %v\n", email.Value())
	}
}
