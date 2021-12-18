package handler

import(
	"log"
	"net/http"
	"io/ioutil"
)

func (h *Handler) ImageStorageHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		/*
		cookie, err := req.Cookie("access-token")
		if err != nil{
			log.Println(err)
			
		}
		userID, err := h.service.ParseToken(cookie.Value)
		if err != nil{
			log.Println(err)
		}
		*/
		

		imgName := req.URL.Query().Get("image")
		imgFile, err := ioutil.ReadFile("img_storage/" + imgName)
		if err != nil {
			resW.WriteHeader(404)
			resW.Write([]byte(`<h1>File not found :(</h1>`))
			log.Println(err)
		}

		resW.Write(imgFile)
	}
}