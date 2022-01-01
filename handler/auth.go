package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"picturestore/entity"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Payload    string `json:"payload"`
}

func returnResponse(resW http.ResponseWriter, res *Response) {
	data, _ := json.Marshal(res)
	resW.Write(data)
}

func (h *Handler) SignUp(resW http.ResponseWriter, req *http.Request) {
	var userdata entity.User
	var response Response

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userdata)
	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusBadRequest)
		response = Response{
			StatusCode: http.StatusBadRequest,
			Payload:    "Invalid data format",
		}
		returnResponse(resW, &response)
		return
	}

	user_id, err := h.service.Auth.SignUp(userdata)

	fmt.Println("NEW USER : ", user_id)

	if err != nil {
		log.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		response = Response{
			StatusCode: http.StatusInternalServerError,
			Payload:    "Internal Server Error",
		}
		returnResponse(resW, &response)
		return
	}
	err = h.service.Storage.NewUserBucket(user_id)
	if err != nil {
		log.Println(err)
		response = Response{
			StatusCode: http.StatusInternalServerError,
			Payload:    "Unable to allocate storage for you",
		}
		returnResponse(resW, &response)
		return
	}
	resW.WriteHeader(http.StatusOK)
	response = Response{
		StatusCode: http.StatusOK,
		Payload:    "Successfully signed up!",
	}

	returnResponse(resW, &response)
}

func (h *Handler) SignIn(resW http.ResponseWriter, req *http.Request) {

	var userdata entity.SignInInput
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userdata)

	fmt.Println("SIGN IN: ", userdata)

	if err != nil {
		log.Fatal(err)
	}

	token, err := h.service.Auth.GenerateToken(userdata)
	if err != nil {
		resW.Write([]byte("Authorization error"))
		log.Fatal(err)
		return
	}

	cookie := http.Cookie{
		Name:  "access-token",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(resW, &cookie)

	http.Redirect(resW, req, "http://localhost:8090/app/my-images", http.StatusSeeOther)

}
