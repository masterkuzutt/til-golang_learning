package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	"til-golang_learning/gin-sample/adapter/controllers"
	"til-golang_learning/gin-sample/adapter/gateway"
	"til-golang_learning/gin-sample/adapter/presenters"
	"til-golang_learning/gin-sample/externals/sqlite"
	"til-golang_learning/gin-sample/usecase"
)

type CustomMessageHandler struct {
	Message string
}

func (m *CustomMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.Message)
}

func CustomMessageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler func yade")
}

func ClosureCustomMessageHandlerFunc(message string) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		})
}

func Login(c *gin.Context) {
	current_user := sessions.Default(c)

	switch c.Request.Method {

	case "GET":
		presenter := presenters.NewLoginPresenter(c)
		presenter.Emit()
	case "POST":
		repository := gateway.NewUserRepository(sqlite.Connect())
		interactor := usecase.NewLoginInteractor(repository)
		controller := controllers.NewLoginController(interactor)

		username := c.PostForm("username")
		password := c.PostForm("password")

		res := controller.Execute(username, password)

		log.Println(res)
		if res.Status == "ok" {
			log.Println("OK")
			current_user.Set("username", username)
			current_user.Save()
			c.Redirect(http.StatusMovedPermanently, "/index.html")

		} else {
			log.Println("NG")
			presenter := presenters.NewLoginPresenter(c)
			presenter.Emit()

		}
		//[TODO]form dataをここで受けてpresenterを読んでるのがすごく気持ち悪いような気もする。presenterを表示だけにするなら
	}
}
func Root(c *gin.Context) {
	//[TODO] ユーザリスト自体不要
	repository := gateway.NewUserRepository(sqlite.Connect())
	interactor := usecase.NewUserListInteractor(repository)
	controller := controllers.NewUserListController(interactor)

	presenter := presenters.NewUserListPresenter(c)

	presenter.Emit(controller.Execute(""))
}

func Register(c *gin.Context) {
	// セッション管理はこっちにあってもいい気がする。
	//presenterには本当にテンプレートのレンダーだけにするか？そうすればJSONにも対応できるのでは。

	// current_user := sessions.Default(c)
	// current_user.Set("is_anonymous", true)
	// current_user.Save()

	switch c.Request.Method {

	case "GET":
		presenter := presenters.NewRegisterPresenter(c)
		presenter.Emit()
	case "POST":
		repository := gateway.NewUserRepository(sqlite.Connect())
		interactor := usecase.NewRegisterInteractor(repository)
		controller := controllers.NewRegisterController(interactor)

		username := c.PostForm("username")
		password := c.PostForm("password")

		res := controller.Execute(username, password)
		if res.Status == "ok" {
			c.Redirect(http.StatusMovedPermanently, "/login.html")
		} else {
			presenter := presenters.NewRegisterPresenter(c)
			presenter.Emit()

		}
		//[TODO]form dataをここで受けてpresenterを読んでるのがすごく気持ち悪いような気もする。presenterを表示だけにするなら
	}
}
