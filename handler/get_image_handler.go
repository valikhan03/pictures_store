package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	cookies := req.Cookies()
	var tokenCookie *http.Cookie

		for _, cookie := range cookies{
			if(cookie.Name == "access-token"){
				tokenCookie = cookie
			}
		}

	userID, err := h.service.ParseToken(tokenCookie.Value)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("user [%s] - GET all files req", userID)

	h.service.GetAllFilesList(userID)	
}
