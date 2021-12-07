package mock

import (
	"fmt"
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
	e := m.funcC()
	return errs.WrapServerError(errs.ErrUserNotFound, e)
}
func (m ErrorModule) funcC() error {
	return m.funcD()
}
func (m ErrorModule) funcD() error {
	return m.FuncE()
}
func (m ErrorModule) FuncE() error {
	//e := errs.ErrAuth
	e := m.funcF()
	return errs.WrapServerError(errs.ErrAuth, e)
}
func (m ErrorModule) funcF() error {
	return m.funcG()
}
func (m ErrorModule) funcG() error {
	return m.FuncH()
}
func (m ErrorModule) FuncH() error {
	_, err := os.Open("dummy.txt")
	return errs.WrapServerError(errs.ErrUnexpected, err)
}

type ErrorModuleTwo struct {
}

func (m ErrorModuleTwo) Main() error {
	e3 := errors.Wrap(m.FuncAB(), fmt.Sprintf("%+v", m.FuncBA()))
	return e3
}
func (m ErrorModuleTwo) FuncAA() error {
	return m.FuncAB()
}
func (m ErrorModuleTwo) FuncAB() error {
	return errors.Wrap(m.FuncAC(), "test1")
}
func (m ErrorModuleTwo) FuncAC() error {
	return m.FuncAD()
}
func (m ErrorModuleTwo) FuncAD() error {
	return errors.New("dummy error root A")
}
func (m ErrorModuleTwo) FuncBA() error {
	return m.FuncBB()
}
func (m ErrorModuleTwo) FuncBB() error {
	return errors.Wrap(m.FuncBC(), "test2")
}
func (m ErrorModuleTwo) FuncBC() error {
	return m.FuncBD()
}
func (m ErrorModuleTwo) FuncBD() error {
	return errors.New("dummy error root B")
}
