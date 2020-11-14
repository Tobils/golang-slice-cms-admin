package main

import (
	"cms/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Dashboard)
	mux.HandleFunc("/buttons", controller.Buttons)
	mux.HandleFunc("/cards", controller.Cards)
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/register", controller.Register)
	mux.HandleFunc("/forgot-password", controller.ForgotPassword)
	mux.HandleFunc("/blank", controller.Blank)
	mux.HandleFunc("/charts", controller.Charts)
	mux.HandleFunc("/tables", controller.Tables)

	// load static file
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting on port 8081")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
