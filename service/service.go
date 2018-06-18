package service

import (
	"github.com/evil-go/fall"
	"github.com/evil-go/evil-sample/dao"
	"github.com/evil-go/evil-sample/model"
)

type GreetService interface {
	Greeting(string) model.GreetResponse
}

type greetServiceImpl struct {
	Dao dao.GreetDao `wire:""`
}

func (ssi greetServiceImpl) Greeting(name string) model.GreetResponse {
	return model.GreetResponse{Message: ssi.Dao.GreetingForName(name)}
}

func init() {
	fall.Register(&greetServiceImpl{})
}
