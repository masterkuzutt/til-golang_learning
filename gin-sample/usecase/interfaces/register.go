package interfaces

import (
	"til-golang_learning/gin-sample/adapter/context"
)

type IRegisterInputPort interface {
	Handle(context.RegisterRequest) context.RegisterResponese
}

// type IUserListOutputPort interface {
// 	Emit([]string)
// }
