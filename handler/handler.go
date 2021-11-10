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
	router.HandleFunc("/resource", h.MyFilesHandler)
	router.HandleFunc("/img", h.ImageStorageHandler)
	router.HandleFunc("/addimg", h.ImageUploadHandler)

	return router
}
