package handlers

import (
	"log"
	"os"
	"path/filepath"

	// "net/http"

	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {

	current_session := sessions.Default(c)
	log.Println("Login?:", current_session.Get("username"))

	tplpath := filepath.Join(os.Getenv("GOPATH"), "src/til-golang_learning/gin-sample/handlers/templates/index.html")

	tpl, err := pongo2.FromFile(tplpath)
	if err != nil {
		c.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath)
		log.Println(err)
	}

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
