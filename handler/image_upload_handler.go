package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) ImageUploadHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseMultipartForm(32 << 20)
		files := req.MultipartForm.File["file-upload"]

		for _, fileheader := range files {
			file, err := fileheader.Open()
			err = h.service.Upload.UploadFile(fileheader.Filename, file)
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
