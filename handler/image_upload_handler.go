package handler

import(
	"log"
	"net/http"
	"html/template"
)

func (h *Handler) ImageUploadHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseMultipartForm(10 << 20)
		file, fileHeader, err := req.FormFile("file-upload")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		err  = h.service.Upload.UploadFile(fileHeader.Filename, file)
		if err != nil{
			log.Fatal(err)
		}
		resW.Write([]byte("file uploaded"))
	}

	if req.Method == http.MethodGet{
		tmp, err := template.ParseFiles("templates/image-upload/image_upload.htm")
		if err != nil {
			log.Println(err)
		}

		tmp.Execute(resW, nil)

	}
}