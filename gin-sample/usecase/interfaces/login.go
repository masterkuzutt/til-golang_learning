package interfaces

import (
	"til-golang_learning/gin-sample/adapter/context"
)

type ILoginInputPort interface {
	Handle(context.LoginRequest) context.LoginResponese
}

// type IUserListOutputPort interface {
// 	Emit([]string)
// }
