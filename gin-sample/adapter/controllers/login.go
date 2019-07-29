package controllers

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/usecase/interfaces"
)

type loginController struct {
	interactor interfaces.ILoginInputPort
}

func NewLoginController(interactor interfaces.ILoginInputPort) *loginController {
	return &loginController{interactor: interactor}
}

func (uc *loginController) Execute(username string, password string) (res context.LoginResponese) {

	res = uc.interactor.Handle(context.LoginRequest{Username: username, Password: password})
	return
}
