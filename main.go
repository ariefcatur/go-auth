package main

import (
	"fmt"
	authcontroller "login-register/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("lalalala server berlariii di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
