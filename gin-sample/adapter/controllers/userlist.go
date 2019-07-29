package controllers

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/usecase/interfaces"
)

type userListController struct {
	interactor interfaces.IUserListInputPort
	// Interactor usecase.UserInteractor
}

func NewUserListController(interactor interfaces.IUserListInputPort) *userListController {
	return &userListController{interactor: interactor}
}

func (uc *userListController) Execute(username string) (res context.UserListResponese) {

	res = uc.interactor.Handle(context.UserListRequest{Username: username})
	return
}
