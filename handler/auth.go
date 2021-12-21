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
			log.Println(err)
			resW.WriteHeader(http.StatusBadRequest)
			return
		}

		user_id, err := h.service.Auth.SignUp(userdata)
		if err != nil{
			log.Println(err)
			resW.WriteHeader(http.StatusInternalServerError)
			resW.Write([]byte("Internal Server Error"))
			return
		}
		err = h.service.Storage.NewUserBucket(user_id)
		if err != nil{
			log.Println(err)
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
			Name: "access-token",
			Value: token,
			Path: "/app",
		}

		http.SetCookie(resW, &cookie)
	}
}