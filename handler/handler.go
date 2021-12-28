package handler

import (
	"picturestore/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/sign-up", h.SignUpPage).Methods("GET")
	router.HandleFunc("/sign-in", h.SignInPage).Methods("GET")

	router.HandleFunc("/sign-up", h.SignUp).Methods("POST")
	router.HandleFunc("/sign-in", h.SignIn).Methods("POST")

	app := router.PathPrefix("/app/").Subrouter()
	app.Use(h.identifyUser)
	app.HandleFunc("/my-images/{img}", h.GetImage).Methods("GET")
	app.HandleFunc("/upload", h.ImageUploadHandler).Methods("POST")
	app.HandleFunc("/my-files", h.GetAllFilesList).Methods("GET")

	app.HandleFunc("/upload", h.ImageUploadPage).Methods("GET")

	return router
}
