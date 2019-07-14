package route

import (
	"log"
	"os"
	"path/filepath"

	"net/http"

	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// get form data
	current_user := sessions.Default(c)
	if current_user.Get("login") == true {
		c.Request.URL.Path = "/index.html"
		c.Redirect(http.StatusFound, "/index.html")
	}

	current_user.Set("name", "kuzu")

	tplpath := filepath.Join(os.Getenv("GOPATH"), "src/til-golang_learning/gin-sample/templates/auth/login.html")
	// tplpath := "templates/index.html"

	tpl, err := pongo2.FromFile(tplpath)
	if err != nil {
		c.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath)
		log.Println(err)
	}

	err = tpl.ExecuteWriter(pongo2.Context{
		"title":        "Home",
		"form":         nil,
		"posts":        nil,
		"current_user": current_user,
	}, c.Writer)

	if err != nil {
		c.String(500, "Internal Server Error")
		log.Println(err)
	}
}

func LoginForm(c *gin.Context) {

	// get form data
	current_session := sessions.Default(c)
	if current_session.Get("login") == true {
		c.Request.URL.Path = "/index.html"
		c.Redirect(http.StatusFound, "/index.html")
	}

	username := c.PostForm("username")
	// password := c.PostForm("password")

	//[TODO] create db access and login

	current_session.Set("username", username)
	current_session.Save()
	
	log.Println(username, current_session.Get("username"))
	c.Redirect(http.StatusFound, "/index.html")

}

func Logout(c *gin.Context) {
}
