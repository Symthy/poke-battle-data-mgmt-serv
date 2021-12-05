package mock

import (
	"os"

	"github.com/pkg/errors"

	"github.com/Symthy/PokeRest/pokeRest/errs"
)

type ErrorModule struct {
}

func (m ErrorModule) FuncA() error {
	return m.funcB()
}
func (m ErrorModule) funcB() error {
	e := errs.ErrUserNotFound
	e.Wrap(m.funcC())
	return e
}
func (m ErrorModule) funcC() error {
	return m.funcD()
}
func (m ErrorModule) funcD() error {
	return m.FuncE()
}
func (m ErrorModule) FuncE() error {
	e := errs.NewServerError(
		errors.New("funcE"),
		errs.Error,
		99,
		"funcE test",
	)
	//e := errs.ErrAuth
	e.Wrap(m.funcF())
	return e
}
func (m ErrorModule) funcF() error {
	return m.funcG()
}
func (m ErrorModule) funcG() error {
	return m.FuncH()
}
func (m ErrorModule) FuncH() error {
	_, err := os.Open("dummy.txt")
	return err
}
