package web

import (
	"github.com/evil-go/fall"
	"github.com/evil-go/outboy"
	"github.com/evil-go/evil-sample/service"
	"net/http"
)

type GreetController struct {
	Service service.GreetService `wire:""`
	Path    string                `value:"controller.path.hello"`
}

func (mc GreetController) GetHello(rw http.ResponseWriter, req *http.Request) {
	result := mc.Service.Greeting(req.URL.Query().Get("name"))
	rw.Write([]byte(result.Message))
}

func (mc GreetController) Init() {
	outboy.Register(mc, map[string]string{
		"GetHello": mc.Path,
	})
}

func init() {
	fall.RegisterPropertiesFile("web.properties")
	fall.Register(&GreetController{})
}
