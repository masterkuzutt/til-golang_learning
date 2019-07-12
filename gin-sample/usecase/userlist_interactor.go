package usecase

import (
	"examples/gin-sample/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	outputport     interfaces.IUserListOutputPort
}

// [TODO]このユースケースが引数にEntityをとるのはまずいかもしれない。
// 後で考える。
// func (i *UserInteractor) Add(u domain.User) (int, error) {
// 	return i.UserRepository.Store(u)
// }

func (i *UserInteractor) Handle(username string) {
	users, err := i.UserRepository.FindByName(username)
	if err != nil {

	}
	data := make([]string, len(users))
	for _, user := range users {
		data = append(data, user.Username)
	}
	i.outputport.Emit(data)
}
