package gateway

import (
	"til-golang_learning/gin-sample/domain"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Store(u domain.User) (id int, err error) {
	user := &User{
		Username: u.Username,
		Password: u.Password,
	}

	if err = r.Conn.Create(user).Error; err != nil {
		return
	}
	return int(user.ID), nil
}

func (r *userRepository) FindByName(username string) (d []domain.User, err error) {
	users := []User{}

	if err = r.Conn.Where("Username = ?", username).First(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]domain.User, n)

	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].Username = users[i].Username
		d[i].Password = users[i].Password
	}
	return

}
