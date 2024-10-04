package main

import (
	"go-sso/controllers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/signin", controllers.Login)
	http.HandleFunc("/callback", controllers.Callback)
	http.ListenAndServe(":3000", nil)
}
