package usecase

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/adapter/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	// UserPresenter  presenter.IUserListOutputPort
}

// [TODO]このユースケースが引数にEntityをとるのはまずいかもしれない。
// 後で考える。
// func (i *UserInteractor) Add(u domain.User) (int, error) {
// 	return i.UserRepository.Store(u)
// }

func (i *UserInteractor) Handle(req context.UserListRequest) (res context.UserListResponese) {
	users, err := i.UserRepository.FindByName(req.Username)
	if err != nil {

	}

	data := make([]string, len(users))
	for _, user := range users {
		data = append(data, user.Username)
	}

	res = context.UserListResponese{Data: data}
	return

	// i.Outputport.Emit(data)
	// presenterはrouterが引き継ぐ
	// i.UserPresenter.Emit(context.UserListResponese{Data: data})

}
