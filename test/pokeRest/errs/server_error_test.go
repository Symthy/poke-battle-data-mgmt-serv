package errs

import (
	"fmt"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/errs"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type ServerErrorTestSuite struct {
	suite.Suite
}

func TestServerErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ServerErrorTestSuite))
}

func (suite *ServerErrorTestSuite) TestServerError() {
	suite.Run("wrap", func() {

		dummyModule := mock.ErrorModule{}
		dummyError := dummyModule.FuncA()
		//err := errs.ErrAuth
		//err.Wrap(dummyError)
		//fmt.Printf("dummy:  %+v\n", dummyError)
		//err.Wrap(errs.ErrUserNotFound)
		//suite.Fail("unmatched name")
		fmt.Printf("actual:\n%s\n", errs.BuildErrorMessage(dummyError))
		fmt.Printf("--------------------\n")
		// if err, ok := dummyError.(stackTracer); ok {
		// 	for _, f := range err.StackTrace() {
		// 		fmt.Printf("%+s:%d\n", f, f)
		// 	}
		// }

		// err, ok := errors.Cause(dummyModule.FuncE()).(stackTracer)
		// if !ok {
		// 	panic("oops, err does not implement stackTracer")
		// }
		// fmt.Printf("%+v", err)

		//zap.Stack("").String
		fmt.Printf("--------------------\n")
	})

	suite.Run("wrap test", func() {
		// m := mock.ErrorModuleTwo{}
		// e := m.Main()
		// fmt.Printf("actual:\n%+v\n", e)

		// _, err := os.Open("dummy.txt")
		// if _, ok := err.(stackTracer); !ok {
		// 	panic("oops, err does not implement stackTracer")
		// }
	})
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
