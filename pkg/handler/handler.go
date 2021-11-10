package handler

import(
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/resource", MyFilesHandler)
	router.HandleFunc("/img", ImageStorageHandler)
	router.HandleFunc("/addimg", ImageUploadHandler)

	return router
}