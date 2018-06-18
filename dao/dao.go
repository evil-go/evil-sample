package dao

import (
	"github.com/evil-go/fall"
)

type GreetDao interface {
	GreetingForName(name string) string
}

type greetDaoImpl struct {
	DefaultMessage string `value:"message.default"`
	BobMessage     string `value:"message.bob"`
	JuliaMessage   string `value:"message.julia"`
}

func (gdi greetDaoImpl) GreetingForName(name string) string {
	switch name {
	case "Bob":
		return gdi.BobMessage
	case "Julia":
		return gdi.JuliaMessage
	default:
		return gdi.DefaultMessage
	}
}

func init() {
	fall.RegisterPropertiesFile("dao.properties")
	fall.Register(&greetDaoImpl{})
}
