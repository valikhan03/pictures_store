package handler

import (
	"encoding/json"
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

type FilesList struct{
	Payload []string `json:"payload"`
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

	results := h.service.GetAllFilesList(userID)

	if len(results) == 0{
		resW.Write([]byte("EMPTY"))
		return
	}

	filesList := &FilesList{
		Payload: results,
	}

	filesListJSON, err := json.Marshal(filesList)
	
	resW.Header().Add("Content-Type", "application/json")
	resW.Write(filesListJSON)
}
