package interfaces

import (
	"examples/gin-sample/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
}
