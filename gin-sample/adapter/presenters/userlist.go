package presenters

import (
	"fmt"
	"log"
	"til-golang_learning/gin-sample/adapter/context"

	"github.com/flosch/pongo2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IUserListOutputPort interface {
	Emit(context.UserListResponese)
}

type toConsolePresenter struct {
}

func (p *toConsolePresenter) Emit(res context.UserListResponese) {
	fmt.Println(res.Data)
}

type userListWeb struct {
	context *gin.Context
}

func NewUserListPresenter(c *gin.Context) *userListWeb {
	return &userListWeb{context: c}
}

func (p *userListWeb) Emit(res context.UserListResponese) {

	current_user := sessions.Default(p.context)
	tplpath := "adapter/presenters/templates/auth/login.html"
	tpl, err := pongo2.FromFile(tplpath)

	if err != nil {
		p.context.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath, err)
	}

	err = tpl.ExecuteWriter(pongo2.Context{
		"title":        "Home",
		"form":         nil,
		"posts":        nil,
		"current_user": current_user,
	}, p.context.Writer)

	if err != nil {
		p.context.String(500, "Internal Server Error")
	}
}
