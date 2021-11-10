package main

import (
	"log"
	"net/http"
	"picturestore/repository"
	"picturestore/service"
	"picturestore/handler"

	_ "github.com/lib/pq"
)

func main() {
	fileServer := http.FileServer(http.Dir("./templates/static/"))
	http.Handle("/templates/static/", http.StripPrefix("/templates/static/", fileServer))

	db, err := repository.NewPostgresDB()
	if err != nil{
		log.Fatal(err)
	}
	
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	server := &http.Server{
		Addr: ":8090",
		Handler: handlers.InitRoutes(),
	}

	server.ListenAndServe()
}






