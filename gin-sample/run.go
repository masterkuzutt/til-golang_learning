package main

import (
	"examples/gin-sample/usecase"
	"log"

	// "net/http"

	"strconv"

	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"

	auth_route "examples/gin-sample/auth/route"
	main_route "examples/gin-sample/main/route"

	"examples/gin-sample/externals/sqlite"
	"examples/gin-sample/adapter/controller"
	"examples/gin-sample/adapter/presenter"
	"examples/gin-sample/usecase"
)

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("before logic!")
		c.Next()
		log.Println("after logic")
	}
}

func abortMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		//ログイン済みかどうかのチェックなど実装。
		if id == 0 {
			c.JSON(400, gin.H{"message": "invalid id"})
			c.Abort()
		}
	}
}

func sampleSessionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		if v := session.Get("count"); v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		log.Println(count)

	}

}

func login(c *gin.Context) {
	log.Println("enter login")

}

func main() {

	//[TODO] think when to put

	router := gin.Default()
	router.Use(sampleMiddleware())
	// router.Use(abortMiddleWare())

	//pongo2では不要
	// router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/examples/gin-sample/templates/*"))

	// cookieの場合
	// store := cookie.NewStore([]byte("scret"))
	store := memstore.NewStore([]byte("scret")) // go get github.com/quasoft/memstore
	router.Use(sessions.Sessions("mysession", store))
	router.Use(sampleSessionMiddleWare())

	router.GET("/", main_route.Root)
	router.GET("/index.html", main_route.Root)

	router.GET("/login", auth_route.Login)
	router.POST("/loginForm", auth_route.LoginForm)

	// router.GET("/login", login)
	// router.POST("/login", login)

	// router.Run()
	repository := &repository.UserRepository{"Conn":sqlite.Connect()}
	outputport := &presenter.ToConsolePresenter{}
	interactor := &usecase.UserInteractor{"UserRepository": ,"outputport": outputportut}
	controller := &controller.UserListController{}
	controller.Execute("kuzu")
}
