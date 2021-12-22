package handler

import (
	"log"
	"net/http"
)

func (h *Handler) ImageUploadHandler(resW http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("access-token")
	if err != nil{
		log.Println(err)
		
	}
	userID, err := h.service.ParseToken(cookie.Value)
	if err != nil{
		log.Println(err)
	}
	req.ParseMultipartForm(32 << 20) //32mb
	files := req.MultipartForm.File["file-upload"]
		for _, fileheader := range files {
		file, err := fileheader.Open()
		size := fileheader.Size
		
		err = h.service.Storage.UploadFile(userID ,fileheader.Filename, file, size)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
	}

	resW.Write([]byte("file uploaded"))
}
