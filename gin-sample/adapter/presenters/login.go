package presenters

import (
	"log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type loginPresenter struct {
	context *gin.Context
}

func NewLoginPresenter(c *gin.Context) *loginPresenter {
	return &loginPresenter{context: c}
}

// presenter for login page that donot depend on controller.
// bacause if login is sccessful, request is redirect to index.html
func (p *loginPresenter) Emit() {

	tplpath := "adapter/presenters/templates/auth/login.html"
	tpl, err := pongo2.FromFile(tplpath)

	set := pongo2.NewSet("our web templates")
	set.Globals["global_variable"] = "test"

	if err != nil {
		p.context.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath, err)
	}

	err = tpl.ExecuteWriter(pongo2.Context{}, p.context.Writer)
	if err != nil {
		p.context.String(500, "Internal Server Error")
	}

}
