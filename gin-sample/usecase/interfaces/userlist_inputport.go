package interfaces

import (
	"til-golang_learning/gin-sample/adapter/context"
)

type IUserListInputPort interface {
	Handle(context.UserListRequest) context.UserListResponese
}

// type IUserListOutputPort interface {
// 	Emit([]string)
// }
