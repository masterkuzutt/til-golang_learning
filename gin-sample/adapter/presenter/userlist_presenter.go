package presenter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"til-golang_learning/gin-sample/adapter/context"

	"github.com/flosch/pongo2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IUserListOutputPort interface {
	Emit(context.UserListResponese)
}

type ToConsolePresenter struct {
}

func (p *ToConsolePresenter) Emit(res context.UserListResponese) {
	fmt.Println(res.Data)
}

type WebPresenter struct {
	c *gin.Context
}

func (p *WebPresenter) Emit(res context.UserListResponese) {

	current_session := sessions.Default(p.c)
	log.Println("Login?:", current_session.Get("username"))

	tplpath := filepath.Join(os.Getenv("GOPATH"), "src/til-golang_learning/gin-sample/handlers/templates/index.html")

	tpl, err := pongo2.FromFile(tplpath)
	if err != nil {
		p.c.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath)
		log.Println(err)
	}

	for _, v := range res.Data {
		log.Println(v)
	}

	err = tpl.ExecuteWriter(pongo2.Context{
		"title":           "Home",
		"form":            nil,
		"posts":           nil,
		"current_session": current_session,
	}, p.c.Writer)

	if err != nil {
		p.c.String(500, "Internal Server Error")
		log.Println(err)
	}

}

func (p *WebPresenter) SetContext(c *gin.Context) {
	p.c = c
}
