package routes

import (
	"net/http"
	"til-golang_learning/gin-sample/externals/handlers"
)

// net/http
func Routes() {
	mux := http.NewServeMux()
	// fs := http.FileServer(http.Dir("public"))
	mh1 := &handlers.CustomMessageHandler{Message: "hogehogeho"}
	mh2 := http.HandlerFunc(handlers.CustomMessageHandlerFunc)
	mux.Handle("/", mh1)
	mux.Handle("/index.html", mh2)
	mux.Handle("/top.html", handlers.ClosureCustomMessageHandlerFunc("hogehogehogehogehoge"))

	http.ListenAndServe(":8080", mux)
}

// Gin
/*
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
*/
