package handler

import(
	//"log"
	"net/http"
	//"github.com/gorilla/sessions"
)

func (h *Handler) identifyUser(next http.Handler) http.Handler{
	return http.HandlerFunc(func(resW http.ResponseWriter, req *http.Request){
		
		cookies := req.Cookies()
		var tokenCookie *http.Cookie

		for _, cookie := range cookies{
			if(cookie.Name == "access-token"){
				tokenCookie = cookie
			}
		}

		var tokenStr string

		if tokenCookie != nil{
			tokenStr = tokenCookie.Value
		}
		
		

		if len(tokenStr) == 0{
			resW.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err := h.service.Auth.ParseToken(tokenStr)
		if err != nil{
			resW.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(resW, req)
	})
}


