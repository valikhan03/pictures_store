package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"picturestore/entity"
)

func (h *Handler) SignUp(resW http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		var userdata entity.User
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&userdata)
		if err != nil{
			log.Fatal(err)
		}

		err = h.service.Auth.SignUp(userdata)
		if err != nil{
			log.Fatal(err)
			resW.WriteHeader(http.StatusInternalServerError)
			resW.Write([]byte("Internal Server Error"))
		}

		resW.WriteHeader(http.StatusOK)
	}
}


func (h *Handler) SignIn(resW http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		var userdata entity.SignInInput
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&userdata)
		if err != nil{
			log.Fatal(err)
		}

		token, err := h.service.Auth.GenerateToken(userdata)
		if err != nil{
			resW.Write([]byte("Authorization error"))
			log.Fatal(err)
			return
		}

		cookie := http.Cookie{
			Name: "token-auth",
			Value: token,
		}

		http.SetCookie(resW, &cookie)

		
	}
}