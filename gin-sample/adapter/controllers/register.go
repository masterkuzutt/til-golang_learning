package controllers

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/usecase/interfaces"
)

type registerController struct {
	interactor interfaces.IRegisterInputPort
	// Interactor usecase.UserInteractor
}

func NewRegisterController(interactor interfaces.IRegisterInputPort) *registerController {
	return &registerController{interactor: interactor}
}

func (uc *registerController) Execute(username string, password string) (res context.RegisterResponese) {

	res = uc.interactor.Handle(context.RegisterRequest{Username: username, Password: password})
	return
}
