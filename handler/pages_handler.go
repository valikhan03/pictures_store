package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) MainPage(resW http.ResponseWriter, req *http.Request) {
	var file = "templates/index/index.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) MyFilesPage(resW http.ResponseWriter, req *http.Request) {
	var file = ""
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ImageUploadPage(resW http.ResponseWriter, req *http.Request) {
	var file = "templates/image-upload/image_upload.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) MyimagesPage(resW http.ResponseWriter, req *http.Request) {
	var file = "templates/my-images/myimages.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignInPage(resW http.ResponseWriter, req *http.Request) {
	var file = "templates/sign-in/sign_in.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignUpPage(resW http.ResponseWriter, req *http.Request) {
	var file = "templates/sign-up/sign_up.htm"
	tmp, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(resW, nil)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		return
	}
}
