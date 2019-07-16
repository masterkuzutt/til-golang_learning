package controller

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/usecase/interfaces"
)

type UserListController struct {
	//初期化の時に渡してやる？
	Interactor interfaces.IUserListInputPort
}

func (c *UserListController) Execute(username string) (res context.UserListResponese) {

	res = c.Interactor.Handle(context.UserListRequest{Username: username})
	return
}
