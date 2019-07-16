package presenter

import (
	"fmt"
	"til-golang_learning/gin-sample/adapter/context"
)

type IUserListOutputPort interface {
	Emit(context.UserListResponese)
}

type ToConsolePresenter struct {
}

func (p *ToConsolePresenter) Emit(res context.UserListResponese) {
	fmt.Println(res.Data)
}
