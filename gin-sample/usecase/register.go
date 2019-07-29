package usecase

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/adapter/interfaces"
	"til-golang_learning/gin-sample/domain"
)

type registerInteractor struct {
	repository interfaces.UserRepository
	// UserPresenter  presenter.IUserListOutputPort
}

func NewRegisterInteractor(repo interfaces.UserRepository) *registerInteractor {
	return &registerInteractor{repository: repo}
}

// [TODO]このユースケースが引数にEntityをとるのはまずいかもしれない。
// 後で考える。
// func (i *UserInteractor) Add(u domain.User) (int, error) {
// 	return i.UserRepository.Store(u)
// }

func (i *registerInteractor) Handle(req context.RegisterRequest) (res context.RegisterResponese) {
	user := domain.User{Username: req.Username, Password: req.Password}
	_, err := i.repository.Store(user)
	if err != nil {

	}

	res = context.RegisterResponese{Status: "ok"}
	return

	// i.Outputport.Emit(data)
	// presenterはrouterが引き継ぐ
	// i.UserPresenter.Emit(context.UserListResponese{Data: data})

}
