package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"til-golang_learning/gin-sample/adapter/controller"
	"til-golang_learning/gin-sample/adapter/presenter"
)

func Root(c *gin.Context) {

	presenter := &presenter.WebPresenter{}
	contorller := &controller.UserListController{Interactor: Interactor}

	current_session := sessions.Default(c)
	log.Println("Login?:", current_session.Get("username"))

	tplpath := filepath.Join(os.Getenv("GOPATH"), "src/til-golang_learning/gin-sample/handlers/templates/index.html")

	tpl, err := pongo2.FromFile(tplpath)
	if err != nil {
		c.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath)
		log.Println(err)
	}

	userlist := controller.Execute("kuzu")

	err = tpl.ExecuteWriter(pongo2.Context{
		"title":           "Home",
		"form":            nil,
		"posts":           nil,
		"current_session": current_session,
	}, c.Writer)

	if err != nil {
		c.String(500, "Internal Server Error")
		log.Println(err)
	}
}
