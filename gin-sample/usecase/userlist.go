package usecase

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/adapter/interfaces"
)

type userlistInteractor struct {
	repository interfaces.UserRepository
	// UserPresenter  presenter.IUserListOutputPort
}

func NewUserListInteractor(repo interfaces.UserRepository) *userlistInteractor {
	return &userlistInteractor{repository: repo}
}
func (i *userlistInteractor) Handle(req context.UserListRequest) (res context.UserListResponese) {
	users, err := i.repository.FindByName(req.Username)
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
