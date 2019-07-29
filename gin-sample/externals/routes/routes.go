package routes

import (
	"til-golang_learning/gin-sample/externals/handlers"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Routes() {
	//[TODO] think when to put

	router := gin.Default()

	store := memstore.NewStore([]byte("scret")) // go get github.com/quasoft/memstore
	router.Use(sessions.Sessions("current_user", store))

	router.GET("/", handlers.Root)
	router.GET("/index.html", handlers.Root)

	router.GET("/register", handlers.Register)
	router.POST("/register", handlers.Register)

	router.GET("/login", handlers.Login)
	router.POST("/loginForm", handlers.Login)

	router.Run()
}
