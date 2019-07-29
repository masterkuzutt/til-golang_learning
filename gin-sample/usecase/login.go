package usecase

import (
	"til-golang_learning/gin-sample/adapter/context"
	"til-golang_learning/gin-sample/adapter/interfaces"
)

type loginInteractor struct {
	repository interfaces.UserRepository
}

func NewLoginInteractor(repo interfaces.UserRepository) *loginInteractor {
	return &loginInteractor{repository: repo}
}

func (i *loginInteractor) Handle(req context.LoginRequest) (res context.LoginResponese) {
	user, err := i.repository.FindByName(req.Username)
	if err != nil {

	}
	if len(user) == 1 && user[0].Password == req.Password {
		res = context.LoginResponese{Status: "ok"}
	} else {
		res = context.LoginResponese{Status: "ng"}
	}
	return

}
