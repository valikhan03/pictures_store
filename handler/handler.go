package handler

import (
	"picturestore/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service      *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/sign-up", h.SignUp)
	router.HandleFunc("/sign-in", h.SignIn)
	app := router.PathPrefix("/app/").Subrouter()
	app.Use(h.identifyUser)
	app.HandleFunc("/resource", h.MyFilesHandler)
	app.HandleFunc("/img", h.ImageStorageHandler)
	app.HandleFunc("/add-img", h.ImageUploadHandler)

	return router
}
