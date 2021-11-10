package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) MyFilesHandler(resW http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("templates/index/index.htm")
	if err != nil {

		log.Println(err)
	}

	tmp.Execute(resW, nil)
}
