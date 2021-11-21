package command

import "github.com/Symthy/PokeRest/pokeRest/common"

const (
	displayNameField = "displayName"
	passwordField    = "password"
	emailFiled       = "email"
)

func requiredFields() map[string]struct{} {
	return map[string]struct{}{
		"id":   struct{}{},
		"name": struct{}{},
		"role": struct{}{},
	}
}

type GetUserCommand struct {
	targetName   string
	isDoFilter   bool
	filterFields map[string]struct{} // Set
}

func NewGetUserCommand(name string) *GetUserCommand {
	return &GetUserCommand{
		targetName:   name,
		isDoFilter:   false,
		filterFields: map[string]struct{}{},
	}
}

func (c GetUserCommand) TargetName() string {
	return c.targetName
}

func (c GetUserCommand) IsDoFilter() bool {
	return c.isDoFilter
}

func (c GetUserCommand) FilterFields() []string {
	return common.Mapkeys(c.filterFields)
}

func (c *GetUserCommand) SetFilterRequiredFields() *GetUserCommand {
	if !c.isDoFilter {
		c.filterFields = requiredFields()
	}
	return c
}

func (c *GetUserCommand) SetFilterDisplayName() *GetUserCommand {
	c.SetFilterRequiredFields()
	c.filterFields[displayNameField] = struct{}{}
	return c
}

func (c *GetUserCommand) SetFilterEmail() *GetUserCommand {
	c.SetFilterRequiredFields()
	c.filterFields[emailFiled] = struct{}{}
	return c
}

func (c *GetUserCommand) SetFilterPassword() *GetUserCommand {
	c.SetFilterRequiredFields()
	c.filterFields[passwordField] = struct{}{}
	return c
}
