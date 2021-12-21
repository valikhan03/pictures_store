package handler

import(
	"log"
	"net/http"
)

func (h *Handler) ImageStorageHandler(resW http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		
		cookie, err := req.Cookie("access-token")
		if err != nil{
			log.Println(err)
			
		}
		userID, err := h.service.ParseToken(cookie.Value)
		if err != nil{
			log.Println(err)
		}
		
		filename := req.URL.Query().Get("image")
		imgFile, err := h.service.GetFile(userID, filename)
		if err != nil{
			log.Println(err)
			resW.WriteHeader(http.StatusInternalServerError)
			return
		}

		resW.Write(imgFile)
	}
}