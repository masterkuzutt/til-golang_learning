package interfaces

// httpに向けてリクエストとレスポンスを送受信する場合
// たぶん使いどころがない
import (
	"til-golang_learning/gin-sample/adapter/context"
)

type IUserListInputPort interface {
	Handle(context.UserListRequest) context.UserListResponese
}

// type IUserListOutputPort interface {
// 	Emit([]string)
// }
