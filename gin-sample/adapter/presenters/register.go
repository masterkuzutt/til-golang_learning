package presenters

import (
	"log"
	"til-golang_learning/gin-sample/adapter/context"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type IRegisterOutputPort interface {
	Emit(context.UserListResponese)
}

type registerPresenter struct {
	context  *gin.Context
	username string
}

func NewRegisterPresenter(c *gin.Context) *registerPresenter {
	return &registerPresenter{context: c}
}

// presenter for register page that donot depend on controller.
// bacause if register is sccessful, request is redirect to loign.html
func (p *registerPresenter) Emit() {

	tplpath := "adapter/presenters/templates/auth/register.html"
	tpl, err := pongo2.FromFile(tplpath)

	if err != nil {
		p.context.String(500, "Internal Server Error")
		log.Println("Failed to load template:", tplpath, err)
	}

	err = tpl.ExecuteWriter(pongo2.Context{}, p.context.Writer)
	if err != nil {
		p.context.String(500, "Internal Server Error")
	}

}
