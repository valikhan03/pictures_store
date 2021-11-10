package main

import (
	"net/http"
	"picturestore/pkg/handler"
)

func main() {
	fileServer := http.FileServer(http.Dir("./templates/static/"))
	http.Handle("/templates/static/", http.StripPrefix("/templates/static/", fileServer))

	server := &http.Server{
		Addr: ":8090",
		Handler: handler.InitRoutes(),
	}

	server.ListenAndServe()
}






