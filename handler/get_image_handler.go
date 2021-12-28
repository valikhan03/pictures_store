package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"html/template"
)

func (h *Handler) GetImage(resW http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("access-token")
	if err != nil {
		log.Println(err)

	}
	userID, err := h.service.ParseToken(cookie.Value)
	if err != nil {
		log.Println(err)
	}
	filename := mux.Vars(req)["img"]
	imgFile, err := h.service.GetFile(userID, filename)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	resW.Write(imgFile)
}

func (h *Handler) GetAllFilesList(resW http.ResponseWriter, req *http.Request) {

	var file = "templates/my-images/myimages.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie, err := req.Cookie("access_token")
	if err != nil {
		log.Println(err)
	}

	userID, err := h.service.ParseToken(cookie.Value)
	if err != nil {
		log.Println(err)
	}

	h.service.GetAllFilesList(userID)


	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}
