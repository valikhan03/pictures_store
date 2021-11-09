package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./templates/static/"))
	http.Handle("/templates/static/", http.StripPrefix("/templates/static/", fileServer))

	http.HandleFunc("/resource", MyFilesHandler)
	http.HandleFunc("/img", ImageStorageHandler)
	http.HandleFunc("/addimg", ImageUploadHandler)
	http.ListenAndServe(":8090", nil)
}

func MyFilesHandler(resW http.ResponseWriter, req *http.Request) {
	var err error
	tmp, err := template.ParseFiles("templates/index.htm")
	if err != nil {

		log.Println(err)
	}

	tmp.Execute(resW, nil)
}

func ImageStorageHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
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

func ImageUploadHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseMultipartForm(10 << 20)
		file, fileHeader, err := req.FormFile("file-upload")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		dst, err := os.Create("img_storage/" + fileHeader.Filename)
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			log.Fatal(err)
		}

		resW.Write([]byte("file uploaded"))
	}

}
