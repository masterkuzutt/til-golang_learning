package route

import (

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"

	"til-golang_learning/gin-sample/externals/sqlite"
	"til-golang_learning/gin-sample/externals/handers"
)

func test(){
	//[TODO] think when to put

	router := gin.Default()
	router.Use(sampleMiddleware())
	router.Use(abortMiddleWare())

	// pongo2では不要
	// router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/til-golang_learning/gin-sample/templates/*"))

	// cookieの場合
	// store := cookie.NewStore([]byte("scret"))
	// memstoreの場合
	store := memstore.NewStore([]byte("scret")) // go get github.com/quasoft/memstore
	router.Use(sessions.Sessions("mysession", store))
	router.Use(sampleSessionMiddleWare())

	router.GET("/", handlers.Root)
	router.GET("/index.html", handlers.Root)

	// router.GET("/login", auth_route.Login)
	// router.POST("/loginForm", auth_route.LoginForm)

	// router.GET("/login", login)
	// router.POST("/login", login)

	router.Run()

}