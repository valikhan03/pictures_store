package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) ImageUploadHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
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
			
			err = h.service.Upload.UploadFile(userID.String() ,fileheader.Filename, file, size)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}

		resW.Write([]byte("file uploaded"))
	}

	if req.Method == http.MethodGet {
		tmp, err := template.ParseFiles("templates/image-upload/image_upload.htm")
		if err != nil {
			log.Println(err)
		}

		tmp.Execute(resW, nil)

	}
}
