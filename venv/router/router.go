package router

import (
	"Weather_Prediction/venv/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() *mux.Router {
	r := mux.NewRouter()                                                                                             //create router object
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "templates/index.html") }) //handel root url

	r.HandleFunc("/fetch", handler.FetchData).Methods("POST")
	r.HandleFunc("/predict", handler.Prediction).Methods("POST")


	return r
}
